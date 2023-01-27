package main

import (
	"context"

	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.uber.org/fx"

	"github.com/AH-dark/fadian-telegram-bot/app"
)

var (
	ctx         = context.Background()
	tracer      = otel.Tracer("main")
	application *fx.App
)

func init() {
	application = app.New(ctx)
}

func main() {
	ctx, span := tracer.Start(ctx, "main")
	defer span.End()

	logrus.WithContext(ctx).Info("Prepared to run application")
	application.Run()
}
