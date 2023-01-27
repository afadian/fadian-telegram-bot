package app

import (
	"go.uber.org/fx"

	// internal
	"github.com/afadian/fadian-telegram-bot/internal/env"
	"github.com/afadian/fadian-telegram-bot/internal/infra"
	"github.com/afadian/fadian-telegram-bot/internal/logger"
	"github.com/afadian/fadian-telegram-bot/model"
)

func Entry() []fx.Option {
	return []fx.Option{
		fx.Provide(env.NewEnvConfig),
		fx.Provide(logger.NewLogger),
		fx.WithLogger(logger.FxLogger),

		infra.Module(),
		model.Module(),
	}
}
