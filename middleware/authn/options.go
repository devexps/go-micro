package authn

import (
	"context"
	"github.com/devexps/go-micro/v2/middleware/authn/engine"
)

// Option is auth option.
type Option func(*options)

type options struct {
	claims engine.Claims
}

// WithClaims with customer claims
func WithClaims(claims engine.Claims) Option {
	return func(o *options) {
		o.claims = claims
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
