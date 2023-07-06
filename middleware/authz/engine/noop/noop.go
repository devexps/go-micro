package noop

import (
	"context"
	"github.com/devexps/go-micro/v2/middleware/authz/engine"
)

type authorizer struct {
}

var _ engine.Authorizer = (*authorizer)(nil)

func (a authorizer) IsAuthorized(ctx context.Context) error {
	return nil
}
