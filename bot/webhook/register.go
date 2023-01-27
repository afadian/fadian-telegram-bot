package webhook

import (
	"context"

	"github.com/sirupsen/logrus"
	"gopkg.in/telebot.v3"
)

func Register(ctx context.Context, bot *telebot.Bot, webhook *telebot.Webhook) error {
	ctx, span := tracer.Start(ctx, "register-webhook")
	defer span.End()

	logger := logrus.WithContext(ctx)

	if err := bot.SetWebhook(webhook); err != nil {
		logger.WithError(err).Fatal("Failed to set webhook")
		return err
	}

	logger.Info("Webhook registered successfully")

	return nil
}
