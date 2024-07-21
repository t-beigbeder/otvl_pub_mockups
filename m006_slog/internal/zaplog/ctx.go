package zaplog

import (
	"context"
	"go.uber.org/zap"
)

type zapKeyType int

var ZapKey zapKeyType = 1

func NewContextWithZap(parent context.Context, logger *zap.Logger) context.Context {
	return context.WithValue(parent, ZapKey, logger)
}

func Logger(ctx context.Context) *zap.Logger {
	return ctx.Value(ZapKey).(*zap.Logger)
}
