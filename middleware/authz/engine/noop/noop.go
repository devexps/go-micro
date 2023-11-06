package noop

import (
	"context"

	"github.com/devexps/go-micro/v2/middleware/authz/engine"
)

type authorizer struct {
}

var _ engine.Authorizer = (*authorizer)(nil)

// NewAuthorizer .
func NewAuthorizer() engine.Authorizer {
	return &authorizer{}
}

// IsAuthorized .
func (a authorizer) IsAuthorized(ctx context.Context) error {
	return nil
}
