package mqtt

import (
	"context"
	"testing"

	"github.com/devexps/go-micro/v2/broker"

	"github.com/stretchr/testify/assert"
)

const (
	LocalEmqxBroker = "tcp://127.0.0.1:1883"
)

func TestNewBroker(t *testing.T) {
	ctx := context.Background()

	b := NewBroker(
		broker.WithContext(ctx),
		broker.WithAddress(LocalEmqxBroker),
	)
	assert.NotNil(t, b)
	_ = b.Init()

	defer b.Disconnect()

	if err := b.Connect(); err != nil {
		t.Logf("cant connect to broker, skip: %v", err)
		t.Skip()
	}
}
