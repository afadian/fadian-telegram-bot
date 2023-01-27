package router

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"gopkg.in/telebot.v3"

	"github.com/afadian/fadian-telegram-bot/internal/env"
)

var tracer = otel.Tracer("bot.router")

func NewBot(ctx context.Context, config *env.Config) (*telebot.Bot, error) {
	ctx, span := tracer.Start(ctx, "new-router")
	defer span.End()

	bot, err := telebot.NewBot(telebot.Settings{
		URL:     config.Telegram.URL,
		Token:   config.Telegram.Token,
		Updates: config.Telegram.Updates,
		Poller: &telebot.LongPoller{
			Timeout: time.Duration(config.Telegram.PollerTimeout) * time.Second,
		},
		Offline: config.Telegram.Offline,
		OnError: func(err error, c telebot.Context) {
			ctx, span := tracer.Start(c.Get("context").(context.Context), "error-handle")
			defer span.End()

			logrus.
				WithContext(ctx).
				WithError(err).
				WithField("telebot_context", c).
				WithField("telebot_message", c.Message()).
				Error("error in telegram handle")
		},
	})
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Fatal("Failed to create bot instance")
		return nil, err
	}

	logrus.WithContext(ctx).Info("Bot instance created successfully")

	return bot, nil
}
