package app

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.uber.org/fx"
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

	return fx.New(opts...)
}
