package main

import (
	"context"
	"flag"

	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.uber.org/fx"

	"github.com/afadian/fadian-telegram-bot/app"
)

var (
	ctx          = context.Background()
	tracer       = otel.Tracer("main")
	application  *fx.App
	forceMigrate bool
)

func init() {
	ctx, span := tracer.Start(ctx, "initialization")
	defer span.End()

	flag.BoolVar(&forceMigrate, "force-migrate", false, "force migrate database")
	flag.Parse()

	application = app.New(ctx, forceMigrate)
}

func main() {
	ctx, span := tracer.Start(ctx, "application")
	defer span.End()

	logrus.WithContext(ctx).Info("Prepared to run application")
	application.Run()
}
