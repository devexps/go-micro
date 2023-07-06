package authn

import (
	"context"
	"github.com/devexps/go-micro/v2/log"
	"github.com/devexps/go-micro/v2/middleware"
	"github.com/devexps/go-micro/v2/middleware/authn/engine"
)

// Server is a server authenticator middleware.
func Server(authenticator engine.Authenticator) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			claims, err := authenticator.Authenticate(ctx, engine.ContextTypeMicro)
			if err != nil {
				return nil, err
			}
			return handler(engine.NewContext(ctx, claims), req)
		}
	}
}

// Client is a client authenticator middleware
func Client(authenticator engine.Authenticator, opts ...Option) middleware.Middleware {
	o := &options{}
	for _, opt := range opts {
		opt(o)
	}
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			var err error
			if ctx, err = authenticator.CreateIdentity(ctx, engine.ContextTypeMicro, o.claims); err != nil {
				log.Errorf("authenticator middleware create token failed: %v", err.Error())
			}
			return handler(ctx, req)
		}
	}
}

// NewContext injects the provided claims in to the parent context.
func NewContext(ctx context.Context, claims engine.Claims) context.Context {
	return engine.NewContext(ctx, claims)
}

// FromContext extracts the claims from the provided context (if any).
func FromContext(ctx context.Context) (claims engine.Claims, ok bool) {
	return engine.FromContext(ctx)
}
