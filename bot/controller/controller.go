package controller

import (
	"go.opentelemetry.io/otel"
	"go.uber.org/fx"

	"github.com/afadian/fadian-telegram-bot/model"
)

var tracer = otel.Tracer("bot.controller")

type Controller struct {
	fx.In

	Setting model.SettingRepository
}
