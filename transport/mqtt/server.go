package mqtt

import (
	"context"
	"fmt"
	"net/url"
	"sync"

	"github.com/devexps/go-micro/v2/transport"
	"github.com/devexps/go-micro/v2/transport/utils"

	"github.com/devexps/go-micro/broker/mqtt/v2"
	"github.com/devexps/go-micro/v2/broker"
)

var (
	_ transport.Server     = (*Server)(nil)
	_ transport.Endpointer = (*Server)(nil)
)

type SubscriberMap map[string]broker.Subscriber

type SubscribeOption struct {
	handler          broker.Handler
	binder           broker.Binder
	subscribeOptions []broker.SubscribeOption
}
type SubscribeOptionMap map[string]*SubscribeOption

type Server struct {
	broker.Broker
	brokerOpts []broker.Option

	subscribers    SubscriberMap
	subscriberOpts SubscribeOptionMap

	sync.RWMutex
	started bool

	baseCtx context.Context
	err     error

	keepAlive       *utils.KeepAliveService
	enableKeepAlive bool
}

// NewServer news server instance with options
func NewServer(opts ...ServerOption) *Server {
	srv := &Server{
		baseCtx:         context.Background(),
		subscribers:     SubscriberMap{},
		subscriberOpts:  SubscribeOptionMap{},
		brokerOpts:      []broker.Option{},
		started:         false,
		keepAlive:       utils.NewKeepAliveService(nil),
		enableKeepAlive: true,
	}
	srv.init(opts...)

	srv.Broker = mqtt.NewBroker(srv.brokerOpts...)

	return srv
}

// Name .
func (s *Server) Name() string {
	return string(KindMQTT)
}

// Endpoint .
func (s *Server) Endpoint() (*url.URL, error) {
	if s.err != nil {
		return nil, s.err
	}

	return s.keepAlive.Endpoint()
}

// Start .
func (s *Server) Start(ctx context.Context) error {
	if s.err != nil {
		return s.err
	}
	if s.started {
		return nil
	}
	s.err = s.Init()
	if s.err != nil {
		LogErrorf("init broker failed: [%s]", s.err.Error())
		return s.err
	}
	s.err = s.Connect()
	if s.err != nil {
		return s.err
	}
	if s.enableKeepAlive {
		go func() {
			_ = s.keepAlive.Start()
		}()
	}
	LogInfof("server listening on: %s", s.Address())

	s.err = s.doRegisterSubscriberMap()
	if s.err != nil {
		return s.err
	}
	s.baseCtx = ctx
	s.started = true

	return nil
}

// Stop .
func (s *Server) Stop(_ context.Context) error {
	LogInfo("server stopping")
	s.started = false
	return s.Disconnect()
}

// RegisterSubscriber .
func (s *Server) RegisterSubscriber(ctx context.Context, topic string, handler broker.Handler, binder broker.Binder, opts ...broker.SubscribeOption) error {
	s.Lock()
	defer s.Unlock()

	opts = append(opts, broker.WithSubscribeContext(ctx))

	if s.started {
		return s.doRegisterSubscriber(topic, handler, binder, opts...)
	} else {
		s.subscriberOpts[topic] = &SubscribeOption{handler: handler, binder: binder, subscribeOptions: opts}
	}
	return nil
}

// RegisterSubscriber .
func RegisterSubscriber[T any](srv *Server, ctx context.Context, topic string, handler func(context.Context, string, broker.Headers, *T) error, opts ...broker.SubscribeOption) error {
	return srv.RegisterSubscriber(ctx,
		topic,
		func(ctx context.Context, event broker.Event) error {
			switch t := event.Message().Body.(type) {
			case *T:
				if err := handler(ctx, event.Topic(), event.Message().Headers, t); err != nil {
					return err
				}
			default:
				return fmt.Errorf("unsupported type: %T", t)
			}
			return nil
		},
		func() broker.Any {
			var t T
			return &t
		},
		opts...,
	)
}

func (s *Server) init(opts ...ServerOption) {
	for _, o := range opts {
		o(s)
	}
}

func (s *Server) doRegisterSubscriber(topic string, handler broker.Handler, binder broker.Binder, opts ...broker.SubscribeOption) error {
	sub, err := s.Subscribe(topic, handler, binder, opts...)
	if err != nil {
		return err
	}
	s.subscribers[topic] = sub

	return nil
}

func (s *Server) doRegisterSubscriberMap() error {
	for topic, opt := range s.subscriberOpts {
		_ = s.doRegisterSubscriber(topic, opt.handler, opt.binder, opt.subscribeOptions...)
	}
	s.subscriberOpts = SubscribeOptionMap{}

	return nil
}
