package router

import (
	"gopkg.in/telebot.v3"

	"github.com/afadian/fadian-telegram-bot/server/controller"
	"github.com/afadian/fadian-telegram-bot/server/middleware"
)

func Register(bot *telebot.Bot) {
	bot.Use(middleware.ContextInject)

	{
		bot.Handle("/start", controller.StartHandler)
	}
}
