package casbin

import (
	"context"
	stdCasbin "github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/devexps/go-micro/v2/middleware/authz/engine"
)

type authorizer struct {
	opts         *options
	enforcer     *stdCasbin.SyncedEnforcer
	wildcardItem string
}

var _ engine.Authorizer = (*authorizer)(nil)

// NewAuthorizer .
func NewAuthorizer(opts ...Option) (engine.Authorizer, error) {
	auth := &authorizer{
		opts:         &options{},
		wildcardItem: "*",
	}
	for _, o := range opts {
		o(auth.opts)
	}
	var err error

	if auth.opts.model == nil {
		auth.opts.model, err = model.NewModelFromString(DefaultRBACModel)
		if err != nil {
			return nil, err
		}
	}
	auth.enforcer, err = stdCasbin.NewSyncedEnforcer(auth.opts.model, auth.opts.policy)
	if err != nil {
		return nil, err
	}
	if auth.opts.watcher != nil {
		_ = auth.enforcer.SetWatcher(auth.opts.watcher)
	}
	return auth, nil
}

// IsAuthorized .
func (a *authorizer) IsAuthorized(ctx context.Context) error {
	claims, ok := engine.FromContext(ctx)
	if !ok {
		return ErrMissingClaims
	}
	if len(claims.GetSubject()) == 0 || len(claims.GetResource()) == 0 || len(claims.GetAction()) == 0 {
		return ErrInvalidClaims
	}
	var (
		allowed bool
		err     error
	)
	if len(claims.GetProject()) > 0 {
		allowed, err = a.enforcer.Enforce(claims.GetSubject(), claims.GetResource(), claims.GetAction(), claims.GetProject())
	} else {
		allowed, err = a.enforcer.Enforce(claims.GetSubject(), claims.GetResource(), claims.GetAction())
	}
	if err != nil {
		//fmt.Println(allowed, err)
		return ErrUnauthorized
	}
	if !allowed {
		return ErrUnauthorized
	}
	return nil
}
