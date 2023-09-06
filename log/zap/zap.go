package zap

import (
	"fmt"

	"go.uber.org/zap"

	"github.com/devexps/go-micro/v2/log"
)

var _ log.Logger = (*Logger)(nil)

type Logger struct {
	log *zap.Logger
}

// NewLogger creates a zap logger
func NewLogger(logger *zap.Logger) *Logger {
	return &Logger{logger}
}

// Log print the kv pairs log.
func (l *Logger) Log(level log.Level, keyvals ...interface{}) error {
	keylen := len(keyvals)
	if keylen == 0 || keylen%2 != 0 {
		l.log.Warn(fmt.Sprint("Keyvalues must appear in pairs: ", keyvals))
		return nil
	}

	data := make([]zap.Field, 0, (keylen/2)+1)
	for i := 0; i < keylen; i += 2 {
		data = append(data, zap.Any(fmt.Sprint(keyvals[i]), keyvals[i+1]))
	}

	switch level {
	case log.LevelDebug:
		l.log.Debug("", data...)
	case log.LevelInfo:
		l.log.Info("", data...)
	case log.LevelWarn:
		l.log.Warn("", data...)
	case log.LevelError:
		l.log.Error("", data...)
	case log.LevelFatal:
		l.log.Fatal("", data...)
	}
	return nil
}

// Sync the logger
func (l *Logger) Sync() error {
	return l.log.Sync()
}

// Close the logger.
func (l *Logger) Close() error {
	return l.Sync()
}
