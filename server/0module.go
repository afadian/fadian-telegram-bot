package server

import (
	"github.com/afadian/fadian-telegram-bot/server/router"
	"github.com/afadian/fadian-telegram-bot/server/webhook"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module(
		"server",
		fx.Provide(router.NewBot),      // return *telebot.Bot
		fx.Provide(webhook.NewWebhook), // return *telebot.Webhook
		fx.Invoke(webhook.Register),
		fx.Invoke(router.Register),
	)
}
