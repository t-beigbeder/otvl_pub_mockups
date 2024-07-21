package zrlog

import (
	"context"
	"github.com/rs/zerolog"
)

type zeroKeyType int

var ZeroKey zeroKeyType = 1

func NewContextWithZero(parent context.Context, logger *zerolog.Logger) context.Context {
	return context.WithValue(parent, ZeroKey, logger)
}

func Logger(ctx context.Context) *zerolog.Logger {
	return ctx.Value(ZeroKey).(*zerolog.Logger)
}
