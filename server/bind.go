package server

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/telebot.v3"
)

func BindWebhook(r *gin.Engine, webhook *telebot.Webhook) {
	r.Any("/webhook", gin.WrapH(webhook))
}
