package random

import (
	"context"
	"math/rand"

	"github.com/devexps/go-micro/v2/selector"
	"github.com/devexps/go-micro/v2/selector/node/direct"
)

const (
	// Name is random balancer name
	Name = "random"
)

var _ selector.Balancer = (*Balancer)(nil) // Name is balancer name

// Option is random builder option.
type Option func(o *options)

// options is random builder options
type options struct{}

// Balancer is a random balancer.
type Balancer struct{}

// New a random selector.
func New(opts ...Option) selector.Selector {
	return NewBuilder(opts...).Build()
}

// Pick is pick a weighted node.
func (p *Balancer) Pick(_ context.Context, nodes []selector.WeightedNode) (selector.WeightedNode, selector.DoneFunc, error) {
	if len(nodes) == 0 {
		return nil, nil, selector.ErrNoAvailable
	}
	cur := rand.Intn(len(nodes))
	selected := nodes[cur]
	d := selected.Pick()
	return selected, d, nil
}

// NewBuilder returns a selector builder with random balancer
func NewBuilder(opts ...Option) selector.Builder {
	var option options
	for _, opt := range opts {
		opt(&option)
	}
	return &selector.DefaultBuilder{
		Balancer: &Builder{},
		Node:     &direct.Builder{},
	}
}

// Builder is random builder
type Builder struct{}

// Build creates Balancer
func (b *Builder) Build() selector.Balancer {
	return &Balancer{}
}
