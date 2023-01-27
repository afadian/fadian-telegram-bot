package app

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.uber.org/fx"

	"github.com/afadian/fadian-telegram-bot/internal/env"
)

var tracer = otel.Tracer("app")

func New(ctx context.Context, forceMigrate bool) *fx.App {
	ctx, span := tracer.Start(ctx, "New")
	defer span.End()

	opts := []fx.Option{
		fx.Supply(
			fx.Annotate(ctx, fx.As(new(context.Context))),
			forceMigrate,
		),
	}
	opts = append(opts, Entry()...)
	opts = append(opts, fx.Invoke(run))

	return fx.New(opts...)
}

func run(ctx context.Context, config *env.Config, engine *gin.Engine, lc fx.Lifecycle) {
	ctx, span := tracer.Start(ctx, "run")
	defer span.End()

	server := &http.Server{
		Addr:    config.System.Listen,
		Handler: engine,
	}

	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				ctx, span := tracer.Start(ctx, "app-start")
				defer span.End()

				go func() {
					ctx, span := tracer.Start(ctx, "server-start")
					defer span.End()

					logrus.WithContext(ctx).Infof("Starting server on %s", config.System.Listen)
					if err := server.ListenAndServeTLS(config.System.CertFile, config.System.KeyFile); err != nil {
						logrus.WithContext(ctx).WithError(err).Fatal("Failed to start server")
					}
				}()

				return nil
			},
			OnStop: func(ctx context.Context) error {
				ctx, span := tracer.Start(ctx, "app-stop")
				defer span.End()

				if err := server.Shutdown(ctx); err != nil {
					logrus.WithContext(ctx).WithError(err).Error("Failed to shutdown server")
					return err
				}

				logrus.WithContext(ctx).Info("Stopping server")
				return nil
			},
		},
	)
}
