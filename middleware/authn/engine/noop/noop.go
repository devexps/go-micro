package noop

import (
	"context"

	"github.com/devexps/go-micro/v2/middleware/authn/engine"
)

type authenticator struct {
}

var _ engine.Authenticator = (*authenticator)(nil)

// NewAuthenticator .
func NewAuthenticator() engine.Authenticator {
	return &authenticator{}
}

// Authenticate .
func (a authenticator) Authenticate(ctx context.Context, ctxType engine.ContextType) (engine.Claims, error) {
	return &Claims{}, nil
}

// CreateIdentityWithContext .
func (a authenticator) CreateIdentityWithContext(ctx context.Context, ctxType engine.ContextType, claims engine.Claims) (context.Context, error) {
	return ctx, nil
}

// CreateIdentity .
func (a authenticator) CreateIdentity(claims engine.Claims) (string, error) {
	return "", nil
}

// Claims is a custom claims object
type Claims struct {
}

// GetSubject returns a subject string
func (c Claims) GetSubject() string {
	return ""
}
