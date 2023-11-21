package sse

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gorilla/mux"

	"github.com/devexps/go-micro/v2/broker"
	"github.com/devexps/go-micro/v2/encoding"
	"github.com/devexps/go-micro/v2/errors"
	"github.com/devexps/go-micro/v2/log"
	"github.com/devexps/go-micro/v2/transport"
)

type Any interface{}
type MessagePayload Any

var (
	_ transport.Server     = (*Server)(nil)
	_ transport.Endpointer = (*Server)(nil)
	_ http.Handler         = (*Server)(nil)
)

type Server struct {
	*http.Server

	lis     net.Listener
	tlsConf *tls.Config

	network string
	address string
	path    string

	timeout time.Duration

	err   error
	codec encoding.Codec

	router      *mux.Router
	strictSlash bool

	headers    map[string]string
	eventTTL   time.Duration
	bufferSize int

	encodeBase64 bool
	splitData    bool
	autoStream   bool
	autoReplay   bool

	subscribeFunc   SubscriberFunction
	unsubscribeFunc SubscriberFunction

	streamMgr *StreamManager
}

// NewServer will create a server and setup defaults
func NewServer(opts ...ServerOption) *Server {
	srv := &Server{
		network:     "tcp",
		address:     ":0",
		timeout:     1 * time.Second,
		router:      mux.NewRouter(),
		strictSlash: true,
		path:        "/",

		bufferSize:   DefaultBufferSize,
		encodeBase64: false,

		autoStream: false,
		autoReplay: true,
		headers:    map[string]string{},

		streamMgr: NewStreamManager(),
	}

	srv.init(opts...)

	srv.err = srv.listen()

	return srv
}

// Name of server
func (s *Server) Name() string {
	return string(KindSSE)
}

// Endpoint return a real address to registry endpoint.
func (s *Server) Endpoint() (*url.URL, error) {
	addr := s.address

	prefix := "http://"
	if s.tlsConf == nil {
		if !strings.HasPrefix(addr, "http://") {
			prefix = "http://"
		}
	} else {
		if !strings.HasPrefix(addr, "https://") {
			prefix = "https://"
		}
	}
	addr = prefix + addr

	var endpoint *url.URL
	endpoint, s.err = url.Parse(addr)
	return endpoint, nil
}

// Start starts the HTTP server.
func (s *Server) Start(ctx context.Context) error {
	if s.err != nil {
		return s.err
	}
	s.BaseContext = func(net.Listener) context.Context {
		return ctx
	}
	log.Infof("[sse] server listening on: %s", s.lis.Addr().String())

	var err error
	if s.tlsConf != nil {
		err = s.ServeTLS(s.lis, "", "")
	} else {
		err = s.Serve(s.lis)
	}
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	s.HandleServeHTTP(s.path)

	return nil
}

// Stop stops the HTTP server.
func (s *Server) Stop(ctx context.Context) error {
	s.streamMgr.Clean()

	log.Info("[sse] server stopping")
	return s.Shutdown(ctx)
}

// Handle registers a new route with a matcher for the URL path.
func (s *Server) Handle(path string, h http.Handler) {
	s.router.Handle(path, h)
}

// HandlePrefix registers a new route with a matcher for the URL path prefix.
func (s *Server) HandlePrefix(prefix string, h http.Handler) {
	s.router.PathPrefix(prefix).Handler(h)
}

// HandleFunc registers a new route with a matcher for the URL path.
func (s *Server) HandleFunc(path string, h http.HandlerFunc) {
	s.router.HandleFunc(path, h)
}

// HandleHeader registers a new route with a matcher for request header values.
func (s *Server) HandleHeader(key, val string, h http.HandlerFunc) {
	s.router.Headers(key, val).Handler(h)
}

// HandleServeHTTP registers a new route with a matcher for the URL path.
func (s *Server) HandleServeHTTP(path string) {
	s.router.HandleFunc(path, s.ServeHTTP)
}

// Publish sends a message to every client in a streamID.
// If the stream's buffer is full, it blocks until the message is sent out to
// all subscribers (but not necessarily arrived the clients), or when the
// stream is closed.
func (s *Server) Publish(_ context.Context, streamId StreamID, event *Event) {
	stream := s.streamMgr.Get(streamId)
	if stream == nil {
		return
	}
	select {
	case <-stream.quit:
	case stream.event <- s.process(event):
	}
}

// TryPublish is the same as Publish except that when the operation would cause
// the call to be blocked, it simply drops the message and returns false.
// Together with a small BufferSize, it can be useful when publishing the
// latest message ASAP is more important than reliable delivery.
func (s *Server) TryPublish(_ context.Context, streamId StreamID, event *Event) bool {
	stream := s.streamMgr.Get(streamId)
	if stream == nil {
		return false
	}
	select {
	case stream.event <- s.process(event):
		return true
	default:
		return false
	}
}

// PublishData will encode message and Publish it to every client in a streamID
func (s *Server) PublishData(ctx context.Context, streamId StreamID, data MessagePayload) error {
	event := &Event{}

	var err error
	event.Data, err = broker.Marshal(s.codec, data)
	if err != nil {
		return err
	}
	s.Publish(ctx, streamId, event)

	return nil
}

// CreateStream will create a new stream and register it
func (s *Server) CreateStream(streamId StreamID) *Stream {
	stream := s.streamMgr.Get(streamId)
	if stream != nil {
		return stream
	}
	stream = s.createStream(streamId)

	s.streamMgr.Add(stream)

	return stream
}

func (s *Server) init(opts ...ServerOption) {
	for _, o := range opts {
		o(s)
	}
	s.router.StrictSlash(s.strictSlash)
	s.router.NotFoundHandler = http.DefaultServeMux
	s.router.MethodNotAllowedHandler = http.DefaultServeMux

	s.Server = &http.Server{
		Handler:   s.router,
		TLSConfig: s.tlsConf,
	}
}

func (s *Server) listen() error {
	if s.lis == nil {
		lis, err := net.Listen(s.network, s.address)
		if err != nil {
			return err
		}
		s.lis = lis
	}
	return nil
}

func (s *Server) process(event *Event) *Event {
	if s.encodeBase64 {
		event.encodeBase64()
	}
	return event
}

func (s *Server) createStream(streamId StreamID) *Stream {
	stream := newStream(streamId, s.bufferSize, s.autoReplay, s.autoStream, s.subscribeFunc, s.unsubscribeFunc)
	stream.run()
	return stream
}
