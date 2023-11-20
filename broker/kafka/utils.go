package kafka

import (
	"github.com/devexps/go-micro/v2/broker"

	kafkaGo "github.com/segmentio/kafka-go"
)

func kafkaHeaderToMap(h []kafkaGo.Header) broker.Headers {
	m := broker.Headers{}
	for _, v := range h {
		m[v.Key] = string(v.Value)
	}
	return m
}
