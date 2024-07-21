package zaplog

import (
	"context"
	"go.uber.org/zap"
)

func ZaplogF1() {
	zap.L().Info("ZaplogF1")
}

func ZaplogCtxF1(ctx context.Context) {
	Logger(ctx).Info("ZaplogCtxF1 zaplog")
}
