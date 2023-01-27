package router

import (
	"go.uber.org/fx"
	"gopkg.in/telebot.v3"

	"github.com/afadian/fadian-telegram-bot/bot/controller"
	"github.com/afadian/fadian-telegram-bot/bot/middleware"
)

type RouterParams struct {
	fx.In

	Bot        *telebot.Bot
	Controller *controller.Controller
	Middleware *middleware.Middleware
}

func Register(params *RouterParams) {
	params.Bot.Use(params.Middleware.ContextInject)

	{
		params.Bot.Handle("/start", params.Controller.StartHandler)
	}
}
