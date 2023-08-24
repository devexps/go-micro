package validate

import (
	"context"
	"errors"
	microError "github.com/devexps/go-micro/v2/errors"
	"github.com/devexps/go-micro/v2/middleware"
	"testing"
)

// protoV implement validate.validator
type protoV struct {
	name    string
	age     int
	isError bool
}

func (v protoV) Validate() error {
	if v.name == "" || v.age < 0 {
		return errors.New("err")
	}
	return nil
}

func TestTable(t *testing.T) {
	var mock middleware.Handler = func(ctx context.Context, req interface{}) (interface{}, error) {
		return nil, nil
	}
	tests := []protoV{
		{"v1", 365, false},
		{"v2", -1, true},
		{"", 365, true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			v := Validator()(mock)
			_, err := v(context.Background(), test)
			if want, have := test.isError, microError.IsBadRequest(err); want != have {
				t.Errorf("fail data %v, want %v, have %v", test, want, have)
			}
		})
	}
}
