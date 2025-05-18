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
	}{
		"trace": {
			method:      "Trace",
			logLevel:    LogLevelTrace,
			shouldWrite: true,
		},
		"debug": {
			method:      "Debug",
			logLevel:    LogLevelTrace,
			shouldWrite: true,
		},
		"info": {
			method:      "Info",
			logLevel:    LogLevelTrace,
			shouldWrite: true,
		},
		"warning": {
			method:      "Warn",
			logLevel:    LogLevelTrace,
			shouldWrite: true,
		},
		"error": {
			method:      "Error",
			logLevel:    LogLevelTrace,
			shouldWrite: true,
		},
		"fatal": {
			method:      "Fatal",
			logLevel:    LogLevelTrace,
			shouldWrite: true,
		},
		"trace on a higher log level": {
			method:      "Trace",
			logLevel:    LogLevelDebug,
			shouldWrite: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			isExitCalled := false
			osExit = func(code int) { isExitCalled = true }
			defer func() { osExit = os.Exit }()

			msg := "test"

			buffer := new(bytes.Buffer)
			log := New(test.logLevel, buffer)

			method := reflect.ValueOf(log).MethodByName(test.method)
			_ = method.Call([]reflect.Value{reflect.ValueOf(msg)})

			if test.shouldWrite {
				assert.Contains(t, buffer.String(), msg, "the log message should be written")
			} else {
				assert.NotContains(t, buffer.String(), msg, "the log message should not be written")
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
