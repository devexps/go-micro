package sse

import (
	"crypto/tls"
	"github.com/devexps/go-micro/v2/encoding"
	"net"
	"time"
)

const DefaultBufferSize = 1024

type ServerOption func(o *Server)

// WithNetwork with server network
func WithNetwork(network string) ServerOption {
	return func(s *Server) {
		s.network = network
	}
}

// WithAddress with server address
func WithAddress(addr string) ServerOption {
	return func(s *Server) {
		s.address = addr
	}
}

// WithTimeout with server timeout
func WithTimeout(timeout time.Duration) ServerOption {
	return func(s *Server) {
		s.timeout = timeout
	}
}

// WithTLSConfig with server TLS config
func WithTLSConfig(c *tls.Config) ServerOption {
	return func(o *Server) {
		o.tlsConf = c
	}
}

// WithListener with server listener
func WithListener(lis net.Listener) ServerOption {
	return func(s *Server) {
		s.lis = lis
	}
}

// WithBufferSize with server buffer size
func WithBufferSize(size int) ServerOption {
	return func(s *Server) {
		s.bufferSize = size
	}
}

// WithCodec with server codec
func WithCodec(c string) ServerOption {
	return func(s *Server) {
		s.codec = encoding.GetCodec(c)
	}
}

// WithEncodeBase64 with server encode base64
func WithEncodeBase64(enable bool) ServerOption {
	return func(s *Server) {
		s.encodeBase64 = enable
	}
}

// WithAutoStream with server auto stream or not
func WithAutoStream(enable bool) ServerOption {
	return func(s *Server) {
		s.autoStream = enable
	}
}

// WithAutoReply with server auto reply or not
func WithAutoReply(enable bool) ServerOption {
	return func(s *Server) {
		s.autoReplay = enable
	}
}

// WithSplitData with server split data or not
func WithSplitData(enable bool) ServerOption {
	return func(s *Server) {
		s.splitData = enable
	}
}

// WithHeaders with server custom headers
func WithHeaders(headers map[string]string) ServerOption {
	return func(s *Server) {
		s.headers = headers
	}
}

// WithSubscriberFunction with server subscriber function
func WithSubscriberFunction(sub SubscriberFunction, unsub SubscriberFunction) ServerOption {
	return func(s *Server) {
		s.subscribeFunc = sub
		s.unsubscribeFunc = unsub
	}
}

// WithEventTTL with server event TTL
func WithEventTTL(timeout time.Duration) ServerOption {
	return func(s *Server) {
		s.eventTTL = timeout
	}
}
