package app

import (
	"go.uber.org/fx"

	"github.com/AH-dark/fadian-telegram-bot/internal/env"
	"github.com/AH-dark/fadian-telegram-bot/internal/infra"
	"github.com/AH-dark/fadian-telegram-bot/internal/logger"
)

func Entry() []fx.Option {
	return []fx.Option{
		fx.Provide(env.NewEnvConfig),
		fx.Provide(logger.NewLogger),
		fx.WithLogger(logger.FxLogger),

		infra.Module(),
	}
}
