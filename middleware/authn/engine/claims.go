package engine

import (
	"golang.org/x/net/context"
)

type ctxKey string

var (
	authnClaimsContextKey = ctxKey("authn-claims")
)

// Claims interface
type Claims interface {
	GetSubject() string
}

// NewContext injects the provided claims in to the parent context.
func NewContext(ctx context.Context, claims Claims) context.Context {
	return context.WithValue(ctx, authnClaimsContextKey, claims)
}

// FromContext extracts the claims from the provided context (if any).
func FromContext(ctx context.Context) (claims Claims, ok bool) {
	claims, ok = ctx.Value(authnClaimsContextKey).(Claims)
	return
}
