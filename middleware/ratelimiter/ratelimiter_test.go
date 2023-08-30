package ratelimiter

import (
	"context"
	"errors"
	"testing"

	"github.com/devexps/go-pkg/v2/ratelimiter"
)

type (
	ratelimiterMock struct {
		reached bool
	}
	ratelimiterReachedMock struct {
		reached bool
	}
)

func (r *ratelimiterMock) Allow() (ratelimiter.DoneFunc, error) {
	return func(_ ratelimiter.DoneInfo) {
		r.reached = true
	}, nil
}

func (r *ratelimiterReachedMock) Allow() (ratelimiter.DoneFunc, error) {
	return func(_ ratelimiter.DoneInfo) {
		r.reached = true
	}, errors.New("errored")
}

func TestServer(t *testing.T) {
	nextValid := func(ctx context.Context, req interface{}) (interface{}, error) {
		return "Hello valid", nil
	}

	rlm := &ratelimiterMock{}
	rlrm := &ratelimiterReachedMock{}

	_, _ = Server(WithLimiter(rlm))(nextValid)(context.Background(), nil)
	if !rlm.reached {
		t.Error("The rate limiter must run the done function.")
	}

	_, _ = Server(WithLimiter(rlrm))(nextValid)(context.Background(), nil)
	if rlrm.reached {
		t.Error("The rate limiter must not run the done function and should be denied.")
	}
}
