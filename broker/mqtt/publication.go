package mqtt

import "github.com/devexps/go-micro/v2/broker"

type publication struct {
	topic string
	msg   *broker.Message
	err   error
}

// Ack .
func (p *publication) Ack() error {
	return nil
}

// Error .
func (p *publication) Error() error {
	return p.err
}

// Topic .
func (p *publication) Topic() string {
	return p.topic
}

// Message .
func (p *publication) Message() *broker.Message {
	return p.msg
}

// RawMessage .
func (p *publication) RawMessage() interface{} {
	return p.msg
}
