package engine

import "context"

// Authorizer interface
type Authorizer interface {
	// IsAuthorized returns nil if everything is ok
	IsAuthorized(ctx context.Context) error
}
