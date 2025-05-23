package app

import (
	"context"

	"github.com/kvitrvn/githeim/internal/config"
)

func Run(ctx context.Context, cfg *config.Config) error {
	return StartServer(ctx, cfg.Server.Port)
}
