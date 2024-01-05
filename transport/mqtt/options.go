package mqtt

import (
	"crypto/tls"

	"github.com/devexps/go-micro/broker/mqtt/v2"
	"github.com/devexps/go-micro/v2/broker"
)

type ServerOption func(o *Server)

// WithBrokerOptions MQ options
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

// WithTLSConfig .
func WithTLSConfig(c *tls.Config) ServerOption {
	return func(s *Server) {
		if c != nil {
			s.brokerOpts = append(s.brokerOpts, broker.WithEnableSecure(true))
		}
		s.brokerOpts = append(s.brokerOpts, broker.WithTLSConfig(c))
	}
}

// WithEnableKeepAlive enable keep alive
func WithEnableKeepAlive(enable bool) ServerOption {
	return func(s *Server) {
		s.enableKeepAlive = enable
	}
}

// WithCleanSession .
func WithCleanSession(enable bool) ServerOption {
	return func(s *Server) {
		s.brokerOpts = append(s.brokerOpts, mqtt.WithCleanSession(enable))
	}
}

// WithAuth .
func WithAuth(username string, password string) ServerOption {
	return func(s *Server) {
		s.brokerOpts = append(s.brokerOpts, mqtt.WithAuth(username, password))
	}
}

// WithClientId .
func WithClientId(clientId string) ServerOption {
	return func(s *Server) {
		s.brokerOpts = append(s.brokerOpts, mqtt.WithClientId(clientId))
	}
}

// WithCodec .
func WithCodec(c string) ServerOption {
	return func(s *Server) {
		s.brokerOpts = append(s.brokerOpts, broker.WithCodec(c))
	}
}
