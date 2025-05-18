package logger

import (
	"bytes"
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogger(t *testing.T) {
	var tests = map[string]struct {
		method      string
		logLevel    LogLevel
		shouldWrite bool
		buffer      *bytes.Buffer
	}{
		"trace": {
			method:      "Trace",
			logLevel:    LogLevelTrace,
			buffer:      new(bytes.Buffer),
			shouldWrite: true,
		},
		"debug": {
			method:      "Debug",
			logLevel:    LogLevelTrace,
			buffer:      new(bytes.Buffer),
			shouldWrite: true,
		},
		"info": {
			method:      "Info",
			logLevel:    LogLevelTrace,
			buffer:      new(bytes.Buffer),
			shouldWrite: true,
		},
		"warning": {
			method:      "Warn",
			logLevel:    LogLevelTrace,
			buffer:      new(bytes.Buffer),
			shouldWrite: true,
		},
		"error": {
			method:      "Error",
			logLevel:    LogLevelTrace,
			buffer:      new(bytes.Buffer),
			shouldWrite: true,
		},
		"fatal": {
			method:      "Fatal",
			logLevel:    LogLevelTrace,
			buffer:      new(bytes.Buffer),
			shouldWrite: true,
		},
		"trace on a higher log level": {
			method:      "Trace",
			logLevel:    LogLevelDebug,
			buffer:      new(bytes.Buffer),
			shouldWrite: false,
		},
		"trace on default output": {
			method:      "Trace",
			logLevel:    LogLevelTrace,
			buffer:      nil,
			shouldWrite: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			isExitCalled := false
			osExit = func(code int) { isExitCalled = true }
			defer func() { osExit = os.Exit }()

			var log Logger

			if test.buffer == nil {
				log = New(test.logLevel, nil)
			} else {
				log = New(test.logLevel, test.buffer)
			}

			msg := "test"

			method := reflect.ValueOf(log).MethodByName(test.method)
			_ = method.Call([]reflect.Value{reflect.ValueOf(msg)})

			if test.shouldWrite {
				assert.Contains(t, test.buffer.String(), msg, "the log message should be written")
			} else if test.buffer != nil {
				assert.NotContains(t, test.buffer.String(), msg, "the log message should not be written")
			}

			// Exit should be called in the Fatal method.
			if test.method == "Fatal" {
				assert.True(t, isExitCalled, "exit should have been called")
			} else {
				assert.False(t, isExitCalled, "exit should not have been called")
			}
		})
	}
}
