package kafka

import (
	"github.com/devexps/go-micro/broker/kafka/v2"
	"github.com/devexps/go-micro/v2/broker"

	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

type ServerOption func(o *Server)

// WithBrokerOptions .
func WithBrokerOptions(opts ...broker.Option) ServerOption {
	return func(s *Server) {
		s.brokerOpts = append(s.brokerOpts, opts...)
	}
}

// WithAddress .
func WithAddress(addrs []string) ServerOption {
	return func(s *Server) {
		s.brokerOpts = append(s.brokerOpts, broker.WithAddress(addrs...))
	}
}

// WithCodec .
func WithCodec(c string) ServerOption {
	return func(s *Server) {
		s.brokerOpts = append(s.brokerOpts, broker.WithCodec(c))
	}
}

// WithEnableKeepAlive enable keep alive
func WithEnableKeepAlive(enable bool) ServerOption {
	return func(s *Server) {
		s.enableKeepAlive = enable
	}
}

// WithPlainMechanism .
func WithPlainMechanism(username, password string) ServerOption {
	return func(s *Server) {
		s.brokerOpts = append(s.brokerOpts, kafka.WithPlainMechanism(username, password))
	}
}

// WithTracerProvider .
func WithTracerProvider(provider trace.TracerProvider, tracerName string) ServerOption {
	return func(s *Server) {
		s.brokerOpts = append(s.brokerOpts, broker.WithTracerProvider(provider, tracerName))
	}
}

// WithPropagator .
func WithPropagator(propagators propagation.TextMapPropagator) ServerOption {
	return func(s *Server) {
		s.brokerOpts = append(s.brokerOpts, broker.WithPropagator(propagators))
	}
}

// WithGlobalTracerProvider .
func WithGlobalTracerProvider() ServerOption {
	return func(s *Server) {
		s.brokerOpts = append(s.brokerOpts, broker.WithGlobalTracerProvider())
	}
}

// WithGlobalPropagator .
func WithGlobalPropagator() ServerOption {
	return func(s *Server) {
		s.brokerOpts = append(s.brokerOpts, broker.WithGlobalPropagator())
	}
}
