package mqtt

import "github.com/devexps/go-micro/v2/log"

///
/// ErrorLogger
///

type ErrorLogger struct{}

// Println .
func (ErrorLogger) Println(v ...interface{}) {
	log.Error(v...)
}

// Printf .
func (ErrorLogger) Printf(format string, v ...interface{}) {
	log.Errorf(format, v...)
}

///
/// CriticalLogger
///

type CriticalLogger struct{}

// Println .
func (CriticalLogger) Println(v ...interface{}) {
	log.Fatal(v...)
}

// Printf .
func (CriticalLogger) Printf(format string, v ...interface{}) {
	log.Fatalf(format, v...)
}

///
/// WarnLogger
///

type WarnLogger struct{}

// Println .
func (WarnLogger) Println(v ...interface{}) {
	log.Warn(v...)
}

// Printf .
func (WarnLogger) Printf(format string, v ...interface{}) {
	log.Warnf(format, v...)
}

///
/// DebugLogger
///

type DebugLogger struct{}

// Println .
func (DebugLogger) Println(v ...interface{}) {
	log.Debug(v...)
}

// Printf .
func (DebugLogger) Printf(format string, v ...interface{}) {
	log.Debugf(format, v...)
}
