package main

import (
	"context"
	"go.uber.org/zap"
	"go.uber.org/zap/exp/zapslog"
	"log/slog"
	"m006_slog/internal/islog"
	"m006_slog/internal/zaplog"
	"m006_slog/pkg/elog"
)

func init() {
	slog.Info("init mainzp")
	zapdl := zap.Must(zap.NewProduction())
	slogl := slog.New(zapslog.NewHandler(zapdl.Core(), nil))
	zap.ReplaceGlobals(zapdl)
	slog.SetDefault(slogl)
}

var zapctx context.Context
var slogctx context.Context

func init() {
	zapctx = zaplog.NewContextWithZap(context.Background(), zap.Must(zap.NewDevelopment(zap.AddCaller())))
	zapdl := zaplog.Logger(zapctx)
	slogctx = elog.NewContextWithSlog(zapctx, slog.New(zapslog.NewHandler(zapdl.Core(), &zapslog.HandlerOptions{AddSource: true})))
}

func main() {
	logger := zap.Must(zap.NewDevelopment())

	defer logger.Sync()

	logger.Info("Hello from Zap logger!")
	sl := slog.New(zapslog.NewHandler(logger.Core(), nil))
	sl.Info("Hello from Zap through slog logger!")
	sub()
	zaplog.ZaplogF1()
	islog.IlogF1()
	zaplog.ZaplogCtxF1(slogctx)
	islog.IlogCtxF1(slogctx)

	sl2 := elog.AsSlog(logger)
	sl2.Info("Hello from Slog sl2")
}

func sub() {
	zap.L().Info("Hello from Zap!")
	slog.Info("Hello from Zap through slog logger!")
}
