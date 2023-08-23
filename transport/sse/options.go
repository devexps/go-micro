package sse

import (
	"crypto/tls"
	"github.com/devexps/go-micro/v2/encoding"
	"net"
	"time"
)

const DefaultBufferSize = 1024

type ServerOption func(o *Server)

func WithNetwork(network string) ServerOption {
	return func(s *Server) {
		s.network = network
	}
}

func WithAddress(addr string) ServerOption {
	return func(s *Server) {
		s.address = addr
	}
}

func WithTimeout(timeout time.Duration) ServerOption {
	return func(s *Server) {
		s.timeout = timeout
	}
}

func WithTLSConfig(c *tls.Config) ServerOption {
	return func(o *Server) {
		o.tlsConf = c
	}
}

func WithListener(lis net.Listener) ServerOption {
	return func(s *Server) {
		s.lis = lis
	}
}

func WithBufferSize(size int) ServerOption {
	return func(s *Server) {
		s.bufferSize = size
	}
}

func WithCodec(c string) ServerOption {
	return func(s *Server) {
		s.codec = encoding.GetCodec(c)
	}
}

func WithEncodeBase64(enable bool) ServerOption {
	return func(s *Server) {
		s.encodeBase64 = enable
	}
}

func WithAutoStream(enable bool) ServerOption {
	return func(s *Server) {
		s.autoStream = enable
	}
}

func WithAutoReply(enable bool) ServerOption {
	return func(s *Server) {
		s.autoReplay = enable
	}
}

func WithSplitData(enable bool) ServerOption {
	return func(s *Server) {
		s.splitData = enable
	}
}

func WithHeaders(headers map[string]string) ServerOption {
	return func(s *Server) {
		s.headers = headers
	}
}

func WithSubscriberFunction(sub SubscriberFunction, unsub SubscriberFunction) ServerOption {
	return func(s *Server) {
		s.subscribeFunc = sub
		s.unsubscribeFunc = unsub
	}
}

func WithEventTTL(timeout time.Duration) ServerOption {
	return func(s *Server) {
		s.eventTTL = timeout
	}
}
