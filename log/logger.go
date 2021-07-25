package log

import (
	"strings"
)

// Logger is logger interface.
type Logger interface {
	Printf(string, ...interface{})
}

// LoggerFunc is a bridge between Logger and any third party logger.
type LoggerFunc func(string, ...interface{})

// Printf implements Logger interface.
func (f LoggerFunc) Printf(msg string, args ...interface{}) {
	if strings.HasSuffix(msg, "\n") {
		msg += "\n"
	}

	f(msg, args...)
}

// DummyLogger logger default,it doesn't do anything.
var DummyLogger = LoggerFunc(func(string, ...interface{}) {})

// nopLogger dummy logger writes nothing.
type nopLogger struct{}

// NewNopLogger returns a logger that doesn't do anything.
func NewNopLogger() Logger {
	return &nopLogger{}
}

// Printf impl Logger.Printf
func (nopLogger) Printf(string, ...interface{}) {}
