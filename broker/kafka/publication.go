package kafka

import (
	"context"
	"github.com/devexps/go-micro/v2/broker"
	kafkaGo "github.com/segmentio/kafka-go"
)

type publication struct {
	topic  string
	err    error
	m      *broker.Message
	ctx    context.Context
	reader *kafkaGo.Reader
	km     kafkaGo.Message
}

// Topic .
func (p *publication) Topic() string {
	return p.topic
}

// Message .
func (p *publication) Message() *broker.Message {
	return p.m
}

// Ack .
func (p *publication) Ack() error {
	return p.reader.CommitMessages(p.ctx, p.km)
}

// Error .
func (p *publication) Error() error {
	return p.err
}
