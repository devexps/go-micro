package kafka

import (
	"sync"

	"github.com/devexps/go-micro/v2/broker"

	kafkaGo "github.com/segmentio/kafka-go"
)

type subscriber struct {
	sync.RWMutex

	k       *kafkaBroker
	topic   string
	options broker.SubscribeOptions
	handler broker.Handler
	reader  *kafkaGo.Reader
	closed  bool
	done    chan struct{}
}

// Options .
func (s *subscriber) Options() broker.SubscribeOptions {
	s.RLock()
	defer s.RUnlock()

	return s.options
}

// Topic .
func (s *subscriber) Topic() string {
	s.RLock()
	defer s.RUnlock()

	return s.topic
}

// Unsubscribe .
func (s *subscriber) Unsubscribe() error {
	s.Lock()
	defer s.Unlock()

	var err error
	if s.reader != nil {
		err = s.reader.Close()
	}
	s.closed = true
	if s.k != nil && s.k.subscribers != nil {
		_ = s.k.subscribers.Remove(s.topic)
	}

	return err
}

// IsClosed .
func (s *subscriber) IsClosed() bool {
	s.RLock()
	defer s.RUnlock()

	return s.closed
}
