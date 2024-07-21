package islog

import (
	"context"
	"log/slog"
	"m006_slog/pkg/elog"
)

func IlogF1() {
	slog.Info("IlogF1")
}

func IlogCtxF1(ctx context.Context) {
	elog.Logger(ctx).Info("IlogCtxF1 elog")
}
