package casbin

import (
	_ "embed"
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
)

//go:embed model/restfull_with_role.conf
var DefaultRestfullWithRoleModel string

// Option is Casbin option
type Option func(*options)

type options struct {
	model   model.Model
	policy  persist.Adapter
	watcher persist.Watcher
}

// WithModel set Model for Casbin
func WithModel(model model.Model) Option {
	return func(o *options) {
		o.model = model
	}
}

// WithPolicy set Adapter for Casbin
func WithPolicy(policy persist.Adapter) Option {
	return func(o *options) {
		o.policy = policy
	}
}

// WithWatcher set Watcher for Casbin
func WithWatcher(watcher persist.Watcher) Option {
	return func(o *options) {
		o.watcher = watcher
	}
}
