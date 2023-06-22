package jwt

import (
	"github.com/devexps/go-micro/v2/errors"
	"github.com/devexps/go-micro/v2/middleware/authn/engine"
)

var (
	ErrMissingBearerToken       = errors.Unauthorized(engine.Reason, "missing bearer token")
	ErrMissingKeyFunc           = errors.Unauthorized(engine.Reason, "missing key func")
	ErrUnauthenticated          = errors.Unauthorized(engine.Reason, "unauthenticated")
	ErrInvalidToken             = errors.Unauthorized(engine.Reason, "invalid token")
	ErrTokenExpired             = errors.Unauthorized(engine.Reason, "token expired")
	ErrUnsupportedSigningMethod = errors.Unauthorized(engine.Reason, "unsupported signing method")
	ErrInvalidClaims            = errors.Unauthorized(engine.Reason, "invalid claims")
	ErrInvalidSubject           = errors.Unauthorized(engine.Reason, "invalid subject")
	ErrGetKeyFailed             = errors.Unauthorized(engine.Reason, "get key failed")
	ErrSignTokenFailed          = errors.Unauthorized(engine.Reason, "sign token failed")
)
