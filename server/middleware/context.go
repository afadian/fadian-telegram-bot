package middleware

import (
	"context"

	"gopkg.in/telebot.v3"
)

func ContextInject(fn telebot.HandlerFunc) telebot.HandlerFunc {
	ctx := context.Background()
	ctx, span := tracer.Start(ctx, "bot-request-inject")
	defer span.End()

	return func(c telebot.Context) error {
		c.Set("context", ctx)
		return fn(c)
	}
}
