package logger

import (
	"context"
	"github.com/afadian/fadian-telegram-bot/internal/env"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"os"
)

var tracer = otel.Tracer("internal.logger")

func NewLogger(ctx context.Context, config *env.Config) *logrus.Logger {
	ctx, span := tracer.Start(ctx, "NewLogger")
	defer span.End()

	logger := logrus.StandardLogger()
	logger.SetFormatter(&logrus.TextFormatter{})
	logger.SetReportCaller(true)
	logger.SetOutput(os.Stdout)

	if config.System.Debug {
		logger.SetLevel(logrus.DebugLevel)
		logger.WithContext(ctx).Debug("debug mode enabled")
	} else {
		logger.SetLevel(logrus.InfoLevel)
	}

	return logger
}
