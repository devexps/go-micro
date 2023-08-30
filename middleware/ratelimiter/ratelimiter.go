package ratelimiter

import (
	"context"

	"github.com/devexps/go-pkg/v2/ratelimiter"
	"github.com/devexps/go-pkg/v2/ratelimiter/lbbr"

	"github.com/devexps/go-micro/v2/errors"
	"github.com/devexps/go-micro/v2/middleware"
)

// ErrLimitExceed is service unavailable due to rate limit exceeded.
var ErrLimitExceed = errors.New(429, "RATELIMIT", "service unavailable due to rate limit exceeded")

// Option is rate limiter option.
type Option func(*options)

// WithLimiter set Limiter implementation, default is L-BBR limiter
func WithLimiter(limiter ratelimiter.RateLimiter) Option {
	return func(o *options) {
		o.limiter = limiter
	}
}

type options struct {
	limiter ratelimiter.RateLimiter
}

// Server rate limiter middleware
func Server(opts ...Option) middleware.Middleware {
	options := &options{
		limiter: lbbr.NewLimiter(),
	}
	for _, o := range opts {
		o(options)
	}
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			done, e := options.limiter.Allow()
			if e != nil {
				// rejected
				return nil, ErrLimitExceed
			}
			// allowed
			reply, err = handler(ctx, req)
			done(ratelimiter.DoneInfo{Err: err})
			return
		}
	}
}
