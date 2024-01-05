package mqtt

import (
	"context"
	"github.com/devexps/go-micro/broker/mqtt/v2"
	"github.com/devexps/go-micro/v2/broker"
	"testing"
)

const (
	LocalEmqxBroker = "tcp://127.0.0.1:1883"
)

func TestServer(t *testing.T) {
	ctx := context.Background()

	srv := NewServer(
		WithAddress([]string{LocalEmqxBroker}),
		WithCodec("json"),
	)
	if err := srv.Start(ctx); err != nil {
		t.Logf("cant start server, skip: %v", err)
		t.Skip()
	}
	defer func() {
		if err := srv.Stop(ctx); err != nil {
			t.Errorf("expected nil got %v", err)
		}
	}()
}

func TestClient(t *testing.T) {
	b := mqtt.NewBroker(
		broker.WithAddress(LocalEmqxBroker),
		broker.WithCodec("json"),
	)
	_ = b.Init()

	if err := b.Connect(); err != nil {
		t.Logf("cant connect to broker, skip: %v", err)
		t.Skip()
	}
	defer b.Disconnect()
}
