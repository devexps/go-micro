package authz

import (
	"context"

	"github.com/devexps/go-micro/v2/middleware"
	"github.com/devexps/go-micro/v2/middleware/authz/engine"
)

// Server is a server authorizer middleware
func Server(authorizer engine.Authorizer) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			err := authorizer.IsAuthorized(ctx)
			if err != nil {
				return nil, err
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
