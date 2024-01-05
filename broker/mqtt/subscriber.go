package mqtt

import (
	"sync"

	"github.com/devexps/go-micro/v2/broker"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

type subscriber struct {
	sync.RWMutex

	options broker.SubscribeOptions
	m       *mqttBroker

	closed bool
	topic  string
	qos    byte

	callback MQTT.MessageHandler
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

	if s.m != nil && s.m.client != nil {
		token := s.m.client.Unsubscribe(s.topic)
		err = token.Error()
	}
	s.closed = true

	if s.m != nil && s.m.subscribers != nil {
		_ = s.m.subscribers.Remove(s.topic)
	}
	return err
}

// IsClosed .
func (s *subscriber) IsClosed() bool {
	s.RLock()
	defer s.RUnlock()

	return s.closed
}
