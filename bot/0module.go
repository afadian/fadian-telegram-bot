package bot

import (
	"go.uber.org/fx"

	"github.com/afadian/fadian-telegram-bot/bot/router"
	"github.com/afadian/fadian-telegram-bot/bot/webhook"
)

func Module() fx.Option {
	return fx.Module(
		"bot",
		fx.Provide(router.NewBot),      // return *telebot.Bot
		fx.Provide(webhook.NewWebhook), // return *telebot.Webhook

		fx.Invoke(webhook.Register),
		fx.Invoke(router.Register),
	)
}
