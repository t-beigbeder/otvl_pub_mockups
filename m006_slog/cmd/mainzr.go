package main

import (
	"context"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	slogzerolog "github.com/samber/slog-zerolog/v2"
	"log/slog"
	"m006_slog/internal/islog"
	"m006_slog/internal/zrlog"
	"m006_slog/pkg/elog"
	"os"
)

var lzl zerolog.Logger

func init() {
	log.Info().Msg("init mainzr")
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.CallerFieldName = "caller"
	lzl = zerolog.New(os.Stdout).With().Caller().Logger()
	slog.SetDefault(slog.New(slogzerolog.Option{Level: slog.LevelDebug, Logger: &lzl}.NewZerologHandler()))
}

var zeroctx context.Context
var slogzrctx context.Context

func init() {
	zl := zerolog.New(os.Stdout).With().Timestamp().Caller().Logger()
	zeroctx = zrlog.NewContextWithZero(context.Background(), &zl)
	slogzrctx = elog.NewContextWithSlog(zeroctx, slog.New(slogzerolog.Option{Level: slog.LevelDebug, Logger: &zl}.NewZerologHandler()))
}

func main() {
	lzl.Info().Msg("Info message")
	lzl.Error().Msg("Error message")
	zrsub()
	islog.IlogF1()
	islog.IlogCtxF1(slogzrctx)
}

func zrsub() {
	lzl.Info().Msg("Hello from Zero!")
	slog.Info("Hello from Zero through slog logger!")
}
