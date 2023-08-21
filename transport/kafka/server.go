package kafka

import (
	"context"
	"github.com/devexps/go-micro/v2/broker"
	"github.com/devexps/go-micro/v2/broker/kafka"
	"github.com/devexps/go-micro/v2/log"
	"github.com/devexps/go-micro/v2/transport"
	"github.com/devexps/go-micro/v2/transport/utils"
	"net/url"
	"sync"
)

var (
	_ transport.Server     = (*Server)(nil)
	_ transport.Endpointer = (*Server)(nil)
)

type SubscriberMap map[string]broker.Subscriber

type SubscribeOption struct {
	handler broker.Handler
	binder  broker.Binder
	opts    []broker.SubscribeOption
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

	keepAlive *utils.KeepAliveService
}

// NewServer creates a server by options
func NewServer(opts ...ServerOption) *Server {
	srv := &Server{
		baseCtx:        context.Background(),
		subscribers:    SubscriberMap{},
		subscriberOpts: SubscribeOptionMap{},
		brokerOpts:     []broker.Option{},
		started:        false,
		keepAlive:      utils.NewKeepAliveService(nil),
	}
	for _, o := range opts {
		o(srv)
	}
	srv.Broker = kafka.NewBroker(srv.brokerOpts...)

	return srv
}

// Endpoint return a real address to registry endpoint.
func (s *Server) Endpoint() (*url.URL, error) {
	if s.err != nil {
		return nil, s.err
	}
	return s.keepAlive.Endpoint()
}

// Start starts a server
func (s *Server) Start(ctx context.Context) error {
	if s.err != nil {
		return s.err
	}
	if s.started {
		return nil
	}
	s.err = s.Init()
	if s.err != nil {
		log.Errorf("[kafka] init broker failed: [%s]", s.err.Error())
		return s.err
	}
	s.err = s.Connect()
	if s.err != nil {
		return s.err
	}
	go func() {
		_ = s.keepAlive.Start()
	}()
	log.Infof("[kafka] server listening on: %s", s.Address())

	s.err = s.doRegisterSubscriberMap()
	if s.err != nil {
		return s.err
	}
	s.baseCtx = ctx
	s.started = true

	return nil
}

// Stop stops a server
func (s *Server) Stop(ctx context.Context) error {
	if s.started == false {
		return nil
	}
	log.Info("[kafka] server stopping")

	for _, v := range s.subscribers {
		if err := v.Unsubscribe(); err != nil {
			log.Errorf("[kafka] un-subscriber (%s) failed: [%s]", v.Topic(), err.Error())
		}
	}
	s.subscribers = SubscriberMap{}
	s.subscriberOpts = SubscribeOptionMap{}

	s.started = false
	return s.Disconnect()
}

// RegisterSubscriber registers a subscriber
// @param ctx is the context
// @param topic is Subscribe topics
// @param queue is Subscribe group
// @param handler is Subscriber handler
func (s *Server) RegisterSubscriber(ctx context.Context, topic, queue string, disableAutoAck bool, handler broker.Handler, binder broker.Binder, opts ...broker.SubscribeOption) error {
	s.Lock()
	defer s.Unlock()

	// opts...
	opts = append(opts, broker.WithQueueName(queue))
	if disableAutoAck {
		opts = append(opts, broker.DisableAutoAck())
	}
	opts = append([]broker.SubscribeOption{broker.WithSubscribeContext(ctx)}, opts...)

	if s.started {
		return s.doRegisterSubscriber(topic, handler, binder, opts...)
	} else {
		s.subscriberOpts[topic] = &SubscribeOption{handler: handler, binder: binder, opts: opts}
	}
	return nil
}

func (s *Server) doRegisterSubscriberMap() error {
	for topic, opt := range s.subscriberOpts {
		if err := s.doRegisterSubscriber(topic, opt.handler, opt.binder, opt.opts...); err != nil {
			log.Errorf("[kafka] do register subscriber (%s) failed: [%s]", topic, err.Error())
		}
	}
	s.subscriberOpts = SubscribeOptionMap{}
	return nil
}

func (s *Server) doRegisterSubscriber(topic string, handler broker.Handler, binder broker.Binder, opts ...broker.SubscribeOption) error {
	sub, err := s.Subscribe(topic, handler, binder, opts...)
	if err != nil {
		return err
	}
	s.subscribers[topic] = sub
	return nil
}
