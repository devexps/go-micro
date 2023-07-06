package casbin

import (
	"github.com/devexps/go-micro/v2/errors"
	"github.com/devexps/go-micro/v2/middleware/authz/engine"
)

var (
	ErrMissingClaims = errors.Forbidden(engine.Reason, "missing claims")
	ErrInvalidClaims = errors.Forbidden(engine.Reason, "invalid claims")
	ErrUnauthorized  = errors.Forbidden(engine.Reason, "unauthorized")
)
