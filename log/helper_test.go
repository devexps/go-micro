package log

import (
	"context"
	"os"
	"testing"
)

func TestHelper(_ *testing.T) {
	logger := With(DefaultLogger, "ts", DefaultTimestamp, "caller", DefaultCaller)
	log := NewHelper(logger)

	log.Log(LevelDebug, "msg", "test debug")
	log.Debug("test debug")
	log.Debugf("test %s", "debug")
	log.Debugw("log", "test debug")

	log.Warn("test warn")
	log.Warnf("test %s", "warn")
	log.Warnw("log", "test warn")
}

func TestHelperWithMsgKey(_ *testing.T) {
	logger := With(DefaultLogger, "ts", DefaultTimestamp, "caller", DefaultCaller)
	log := NewHelper(logger, WithMessageKey("message"))
	log.Debugf("test %s", "debug")
	log.Debugw("log", "test debug")
}

func TestHelperLevel(_ *testing.T) {
	log := NewHelper(DefaultLogger)
	log.Debug("test debug")
	log.Info("test info")
	log.Infof("test %s", "info")
	log.Warn("test warn")
	log.Error("test error")
	log.Errorf("test %s", "error")
	log.Errorw("log", "test error")
}

type traceKey struct{}

func TestContext(_ *testing.T) {
	logger := With(NewStdLogger(os.Stdout),
		"trace", Trace(),
	)
	log := NewHelper(logger)
	ctx := context.WithValue(context.Background(), traceKey{}, "2233")
	log.WithContext(ctx).Info("got trace!")
}

func Trace() Valuer {
	return func(ctx context.Context) interface{} {
		s, ok := ctx.Value(traceKey{}).(string)
		if !ok {
			return nil
		}
		return s
	}
}
