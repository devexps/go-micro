package circuitbreaker

import (
	"context"

	"github.com/devexps/go-micro/v2/errors"
	"github.com/devexps/go-micro/v2/internal/group"
	"github.com/devexps/go-micro/v2/middleware"
	"github.com/devexps/go-micro/v2/middleware/circuitbreaker/breaker"
	"github.com/devexps/go-micro/v2/middleware/circuitbreaker/breaker/sre"
	"github.com/devexps/go-micro/v2/transport"
)

// ErrNotAllowed is request failed due to circuit breaker triggered.
var ErrNotAllowed = errors.New(503, "CIRCUITBREAKER", "request failed due to circuit breaker triggered")

// Option is circuit breaker option.
type Option func(*options)

// WithCircuitBreaker with circuit breaker genFunc.
func WithCircuitBreaker(genBreakerFunc func() breaker.CircuitBreaker) Option {
	return func(o *options) {
		o.group = group.NewGroup(func() interface{} {
			return genBreakerFunc()
		})
	}
}

type options struct {
	group *group.Group
}

// Client will return errBreakerTriggered when the circuit breaker is triggered and the request is rejected directly.
func Client(opts ...Option) middleware.Middleware {
	opt := &options{
		group: group.NewGroup(func() interface{} {
			return sre.NewBreaker()
		}),
	}
	for _, o := range opts {
		o(opt)
	}
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			info, _ := transport.FromClientContext(ctx)
			breaker := opt.group.Get(info.Operation()).(breaker.CircuitBreaker)
			if err := breaker.Allow(); err != nil {
				// rejected
				// NOTE: when client reject requests locally,
				// continue to add counter let the drop ratio higher.
				breaker.MarkFailed()
				return nil, ErrNotAllowed
			}
			// allowed
			reply, err := handler(ctx, req)
			if err != nil && (errors.IsInternalServer(err) || errors.IsServiceUnavailable(err) || errors.IsGatewayTimeout(err)) {
				breaker.MarkFailed()
			} else {
				breaker.MarkSuccess()
			}
			return reply, err
		}
	}
}
