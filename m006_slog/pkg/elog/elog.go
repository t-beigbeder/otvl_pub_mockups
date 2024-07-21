package elog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/exp/zapslog"
	"log/slog"
)

func AsSlog(zapLogger *zap.Logger) *slog.Logger {
	return slog.New(zapslog.NewHandler(zapLogger.Core(), &zapslog.HandlerOptions{AddSource: true}))
}

type SLogger interface {
	AsSlog(logger interface{}) *slog.Logger
}
