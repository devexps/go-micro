package noop

import (
	"context"
	"github.com/devexps/go-micro/v2/middleware/authn/engine"
)

type authenticator struct {
}

var _ engine.Authenticator = (*authenticator)(nil)

// Authenticate .
func (a authenticator) Authenticate(ctx context.Context, ctxType engine.ContextType) (engine.Claims, error) {
	return &Claims{}, nil
}

// CreateIdentity .
func (a authenticator) CreateIdentity(ctx context.Context, ctxType engine.ContextType, claims engine.Claims) (context.Context, error) {
	return ctx, nil
}

// Claims is a custom claims object
type Claims struct {
}

// GetSubject returns a subject string
func (c Claims) GetSubject() string {
	return ""
}
