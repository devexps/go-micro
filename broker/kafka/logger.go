package kafka

import "github.com/devexps/go-micro/v2/log"

type Logger struct {
}

// Printf .
func (l Logger) Printf(msg string, args ...interface{}) {
	log.Infof(msg, args...)
}

type ErrorLogger struct {
}

// Printf .
func (l ErrorLogger) Printf(msg string, args ...interface{}) {
	log.Errorf(msg, args...)
}
