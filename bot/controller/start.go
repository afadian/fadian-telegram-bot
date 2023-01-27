package controller

import (
	"context"

	"github.com/sirupsen/logrus"
	"gopkg.in/telebot.v3"
)

func (ctr *Controller) StartHandler(c telebot.Context) error {
	ctx, span := tracer.Start(c.Get("context").(context.Context), "start-handler")
	defer span.End()

	logger := logrus.WithContext(ctx)

	if err := c.Send("Hi, I'm Fadian Bot. I'm here to help you to manage your Fadian behavior."); err != nil {
		logger.WithError(err).Error("Failed to send message")
		return err
	}

	return nil
}
