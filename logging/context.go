package logging

import (
	"context"

	"github.com/sirupsen/logrus"
)

type loggerMarker string

type logContext struct {
	*logrus.Entry
}

const loggerContextKey = loggerMarker("logger")

func logRPC(ctx context.Context, entry *logrus.Entry) context.Context {
	l := &logContext{entry}
	return context.WithValue(ctx, loggerContextKey, l)
}
