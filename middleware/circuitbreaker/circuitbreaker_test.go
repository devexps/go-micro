package circuitbreaker

import (
	"context"
	"errors"
	"testing"

	goMicroError "github.com/devexps/go-micro/v2/errors"
	"github.com/devexps/go-micro/v2/transport"

	"github.com/devexps/go-pkg/v2/circuitbreaker"
)

type circuitBreakerMock struct {
	err error
}

func (c *circuitBreakerMock) Allow() error { return c.err }
func (c *circuitBreakerMock) MarkSuccess() {}
func (c *circuitBreakerMock) MarkFailed()  {}

type transportMock struct {
	kind      transport.Kind
	endpoint  string
	operation string
}

func (tr *transportMock) Kind() transport.Kind { return tr.kind }

func (tr *transportMock) Endpoint() string { return tr.endpoint }

func (tr *transportMock) Operation() string { return tr.operation }

func (tr *transportMock) RequestHeader() transport.Header { return nil }

func (tr *transportMock) ReplyHeader() transport.Header { return nil }

func TestServer(_ *testing.T) {
	nextValid := func(ctx context.Context, req interface{}) (interface{}, error) {
		return "Hello valid", nil
	}
	nextInvalid := func(ctx context.Context, req interface{}) (interface{}, error) {
		return nil, goMicroError.InternalServer("test", "test internal server")
	}

	ctx := transport.NewClientContext(context.Background(), &transportMock{})

	_, _ = Client(WithCircuitBreaker(func() circuitbreaker.CircuitBreaker {
		return &circuitBreakerMock{err: errors.New("circuitbreaker error")}
	}))(nextValid)(ctx, nil)

	_, _ = Client(func(_ *options) {})(nextValid)(ctx, nil)

	_, _ = Client(func(_ *options) {})(nextInvalid)(ctx, nil)
}
