package webhook

import (
	"context"

	"go.opentelemetry.io/otel"
	"gopkg.in/telebot.v3"

	"github.com/afadian/fadian-telegram-bot/internal/env"
	"github.com/afadian/fadian-telegram-bot/pkg/util"
)

var tracer = otel.Tracer("bot.webhook")

func NewWebhook(ctx context.Context, config *env.Config) *telebot.Webhook {
	ctx, span := tracer.Start(ctx, "new-webhook")
	defer span.End()

	webhook := &telebot.Webhook{
		Listen:         config.System.Listen,
		MaxConnections: 1000,
		SecretToken:    config.System.Secret,
		Endpoint: &telebot.WebhookEndpoint{
			PublicURL: config.System.PublicURL,
		},
		HasCustomCert: true,
		TLS: &telebot.WebhookTLS{
			Cert: util.AbsolutePath(config.System.CertFile),
			Key:  util.AbsolutePath(config.System.KeyFile),
		},
	}

	return webhook
}
