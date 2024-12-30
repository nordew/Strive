package logger

import (
	"log/slog"
	"os"
)

type Logger interface {
	Infof(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
}

type logger struct {
	slog *slog.Logger
}

// New creates a new logger instance using slog
func New() Logger {
	handler := slog.NewTextHandler(os.Stdout, nil)
	slogInstance := slog.New(handler)

	return &logger{slog: slogInstance}
}

func (l *logger) Infof(format string, args ...interface{}) {
	l.slog.Info(format, "args", args)
}

func (l *logger) Debugf(format string, args ...interface{}) {
	l.slog.Debug(format, "args", args)
}

func (l *logger) Errorf(format string, args ...interface{}) {
	l.slog.Error(format, "args", args)
}
