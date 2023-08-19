package kafka

import (
	"github.com/devexps/go-micro/v2/broker"
	kafkaGo "github.com/segmentio/kafka-go"
	"sync"
)

type subscriber struct {
	topic   string
	opts    broker.SubscribeOptions
	handler broker.Handler
	reader  *kafkaGo.Reader
	closed  bool
	done    chan struct{}
	sync.RWMutex
}

// Options .
func (s *subscriber) Options() broker.SubscribeOptions {
	return s.opts
}

// Topic .
func (s *subscriber) Topic() string {
	return s.topic
}

// Unsubscribe .
func (s *subscriber) Unsubscribe() error {
	var err error
	s.Lock()
	defer s.Unlock()
	s.closed = true
	return err
}
