package log

import (
	"bytes"
	"context"
	"strings"
	"testing"
)

func TestFilterAll(_ *testing.T) {
	logger := With(DefaultLogger, "ts", DefaultTimestamp, "caller", DefaultCaller)
	log := NewHelper(NewFilter(logger,
		FilterLevel(LevelDebug),
		FilterKey("username"),
		FilterValue("hello"),
		FilterFunc(testFilterFunc),
	))
	log.Log(LevelDebug, "msg", "test debug")
	log.Info("hello")
	log.Infow("password", "123456")
	log.Infow("username", "go-micro")
	log.Warn("warn log")
}

func TestFilterLevel(_ *testing.T) {
	logger := With(DefaultLogger, "ts", DefaultTimestamp, "caller", DefaultCaller)
	log := NewHelper(NewFilter(NewFilter(logger, FilterLevel(LevelWarn))))
	log.Log(LevelDebug, "msg1", "te1st debug")
	log.Debug("test debug")
	log.Debugf("test %s", "debug")
	log.Debugw("log", "test debug")
	log.Warn("warn log")
}

func TestFilterCaller(_ *testing.T) {
	logger := With(DefaultLogger, "ts", DefaultTimestamp, "caller", DefaultCaller)
	log := NewFilter(logger)
	_ = log.Log(LevelDebug, "msg1", "te1st debug")
	logHelper := NewHelper(NewFilter(logger))
	logHelper.Log(LevelDebug, "msg1", "te1st debug")
}

func TestFilterKey(_ *testing.T) {
	logger := With(DefaultLogger, "ts", DefaultTimestamp, "caller", DefaultCaller)
	log := NewHelper(NewFilter(logger, FilterKey("password")))
	log.Debugw("password", "123456")
}

func TestFilterValue(_ *testing.T) {
	logger := With(DefaultLogger, "ts", DefaultTimestamp, "caller", DefaultCaller)
	log := NewHelper(NewFilter(logger, FilterValue("debug")))
	log.Debugf("test %s", "debug")
}

func TestFilterFunc(_ *testing.T) {
	logger := With(DefaultLogger, "ts", DefaultTimestamp, "caller", DefaultCaller)
	log := NewHelper(NewFilter(logger, FilterFunc(testFilterFunc)))
	log.Debug("debug level")
	log.Infow("password", "123456")
}

func TestFilterFuncWitchLoggerPrefix(t *testing.T) {
	buf := new(bytes.Buffer)
	tests := []struct {
		logger Logger
		want   string
	}{
		{
			logger: NewFilter(With(NewStdLogger(buf), "caller", "caller", "prefix", "whaterver"), FilterFunc(testFilterFuncWithLoggerPrefix)),
			want:   "",
		},
		{
			// Filtered value
			logger: NewFilter(With(NewStdLogger(buf), "caller", "caller"), FilterFunc(testFilterFuncWithLoggerPrefix)),
			want:   "INFO caller=caller msg=msg filtered=***\n",
		},
		{
			// NO prefix
			logger: NewFilter(With(NewStdLogger(buf)), FilterFunc(testFilterFuncWithLoggerPrefix)),
			want:   "INFO msg=msg filtered=***\n",
		},
	}

	for _, tt := range tests {
		err := tt.logger.Log(LevelInfo, "msg", "msg", "filtered", "true")
		if err != nil {
			t.Fatal("err should be nil")
		}
		got := buf.String()
		if got != tt.want {
			t.Fatalf("filter should catch prefix, want %s, got %s.", tt.want, got)
		}
		buf.Reset()
	}
}

func testFilterFunc(level Level, keyvals ...interface{}) bool {
	if level == LevelWarn {
		return true
	}
	for i := 0; i < len(keyvals); i++ {
		if keyvals[i] == "password" {
			keyvals[i+1] = fuzzyStr
		}
	}
	return false
}

func testFilterFuncWithLoggerPrefix(level Level, keyvals ...interface{}) bool {
	if level == LevelWarn {
		return true
	}
	for i := 0; i < len(keyvals); i += 2 {
		if keyvals[i] == "prefix" {
			return true
		}
		if keyvals[i] == "filtered" {
			keyvals[i+1] = fuzzyStr
		}
	}
	return false
}

func TestFilterWithContext(t *testing.T) {
	ctxKey := struct{}{}
	ctxValue := "filter test value"

	v1 := func() Valuer {
		return func(ctx context.Context) interface{} {
			return ctx.Value(ctxKey)
		}
	}

	info := &bytes.Buffer{}

	logger := With(NewStdLogger(info), "request_id", v1())
	filter := NewFilter(logger, FilterLevel(LevelError))

	ctx := context.WithValue(context.Background(), ctxKey, ctxValue)

	_ = WithContext(ctx, filter).Log(LevelInfo, "kind", "test")

	if info.String() != "" {
		t.Error("filter is not woring")
		return
	}

	_ = WithContext(ctx, filter).Log(LevelError, "kind", "test")
	if !strings.Contains(info.String(), ctxValue) {
		t.Error("don't read ctx value")
	}
}
