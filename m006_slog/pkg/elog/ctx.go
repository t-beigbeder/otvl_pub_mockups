package elog

import (
	"context"
	"log/slog"
)

type slogKeyType int

var SlogKey slogKeyType = 1

func NewContextWithSlog(parent context.Context, logger *slog.Logger) context.Context {
	return context.WithValue(parent, SlogKey, logger)
}

func Logger(ctx context.Context) *slog.Logger {
	return ctx.Value(SlogKey).(*slog.Logger)
}
