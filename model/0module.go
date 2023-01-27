package model

import (
	"go.opentelemetry.io/otel"
	"go.uber.org/fx"
)

var tracer = otel.Tracer("model")

func Module() fx.Option {
	return fx.Module(
		"model",
		fx.Provide(NewSettingService),
	)
}
