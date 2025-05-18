package logger

import (
	"fmt"
	"io"
	"os"
	"time"
)

type LogLevel int

const (
	LogLevelTrace = iota
	LogLevelDebug
	LogLevelInfo
	LogLevelWarn
	LogLevelError
	LogLevelFatal
)

type Logger interface {
	log(severity LogLevel, msg string, args ...any) (n int, err error)
	Trace(msg string, args ...any) (n int, err error)
	Debug(msg string, args ...any) (n int, err error)
	Info(msg string, args ...any) (n int, err error)
	Warn(msg string, args ...any) (n int, err error)
	Error(msg string, args ...any) (n int, err error)
	Fatal(msg string, args ...any)
}

type logger struct {
	Logger
	level  LogLevel
	output io.Writer
}

func New(level LogLevel, output io.Writer) Logger {
	if output == nil {
		output = os.Stdout
	}

	return &logger{
		Logger: nil,
		level:  level,
		output: output,
	}
}

func (l *logger) log(severity LogLevel, msg string, args ...any) (n int, err error) {
	if severity < l.level {
		return 0, nil
	}

	timestamp := time.Now().Format(time.Stamp)
	msgFmt := fmt.Sprintf("%s · %s\n", timestamp, msg)

	return fmt.Fprintf(l.output, msgFmt, args...)
}

func (l *logger) Trace(msg string, args ...any) (n int, err error) {
	return l.log(LogLevelTrace, msg, args...)
}

func (l *logger) Debug(msg string, args ...any) (n int, err error) {
	return l.log(LogLevelDebug, msg, args...)
}

func (l *logger) Info(msg string, args ...any) (n int, err error) {
	return l.log(LogLevelInfo, msg, args...)
}

func (l *logger) Warn(msg string, args ...any) (n int, err error) {
	return l.log(LogLevelWarn, msg, args...)
}

func (l *logger) Error(msg string, args ...any) (n int, err error) {
	return l.log(LogLevelError, msg, args...)
}

var osExit = os.Exit

func (l *logger) Fatal(msg string, args ...any) {
	_, _ = l.log(LogLevelFatal, msg, args...)
	osExit(1)
}
