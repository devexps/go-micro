package kafka

import (
	"context"
	"fmt"
	"net/url"
	"sync"

	"github.com/devexps/go-micro/broker/kafka/v2"

	"github.com/devexps/go-micro/v2/broker"
	"github.com/devexps/go-micro/v2/log"
	"github.com/devexps/go-micro/v2/transport"
	"github.com/devexps/go-micro/v2/transport/utils"
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

	srv.doInjectOptions(opts...)

	srv.Broker = kafka.NewBroker(srv.brokerOpts...)

	return srv
}

func (s *Server) doInjectOptions(opts ...ServerOption) {
	for _, o := range opts {
		o(s)
	}
}

func (s *Server) Name() string {
	return string(KindKafka)
}

func (s *Server) Endpoint() (*url.URL, error) {
	if s.err != nil {
		return nil, s.err
	}
	return s.keepAlive.Endpoint()
}

func (s *Server) Start(ctx context.Context) error {
	if s.err != nil {
		return s.err
	}

	if s.started {
		return nil
	}

	s.err = s.Init()
	if s.err != nil {
		log.Errorf("init broker failed: [%s]", s.err.Error())
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

	log.Infof("server listening on: %s", s.Address())

	s.err = s.doRegisterSubscriberMap()
	if s.err != nil {
		return s.err
	}

	s.baseCtx = ctx
	s.started = true

	return nil
}

func (s *Server) Stop(_ context.Context) error {
	if s.started == false {
		return nil
	}
	log.Info("server stopping")

	for _, v := range s.subscribers {
		_ = v.Unsubscribe()
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
		s.subscriberOpts[topic] = &SubscribeOption{handler: handler, binder: binder, subscribeOptions: opts}
	}
	return nil
}

func RegisterSubscriber[T any](srv *Server, ctx context.Context, topic, queue string, disableAutoAck bool, handler func(context.Context, string, broker.Headers, *T) error, opts ...broker.SubscribeOption) error {
	return srv.RegisterSubscriber(ctx,
		topic,
		queue,
		disableAutoAck,
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
