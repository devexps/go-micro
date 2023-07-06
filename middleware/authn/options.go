package authn

import (
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
