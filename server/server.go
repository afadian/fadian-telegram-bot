package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"

	"github.com/afadian/fadian-telegram-bot/internal/env"
)

var tracer = otel.Tracer("server")

func NewServer(ctx context.Context, config *env.Config) *gin.Engine {
	ctx, span := tracer.Start(ctx, "new-server")
	defer span.End()

	r := gin.Default()

	if config.System.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	r.Use(otelgin.Middleware("webhook-server"))

	return r
}
