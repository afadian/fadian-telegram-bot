package infra

import (
	"go.opentelemetry.io/otel"
	"go.uber.org/fx"
)

var tracer = otel.Tracer("internal.infra")

func Module() fx.Option {
	return fx.Module(
		"infra",
		fx.Invoke(Trace),
	)
}
