package log

import (
	"context"
	"fmt"
	"log"
	"os"
)

type Logger interface {
	Info(message string)
	Error(message string)
}

func New() Logger {
	return &logger{
		inner: log.New(os.Stdout, "", log.LstdFlags),
	}
}

type loggerKey struct{}

func WithContext(ctx context.Context, logger Logger) context.Context {
	return context.WithValue(ctx, loggerKey{}, logger)
}

func FromContext(ctx context.Context) Logger {
	return ctx.Value(loggerKey{}).(Logger)
}

var _ Logger = (*logger)(nil)

type logger struct {
	inner *log.Logger
}

func (l *logger) Info(message string) {
	l.inner.Println(fmt.Sprintf("[INFO] %s", message))
}

func (l *logger) Error(message string) {
	l.inner.Println(fmt.Sprintf("[ERROR] %s", message))
}
