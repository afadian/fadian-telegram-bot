package controller

import "go.opentelemetry.io/otel"

var tracer = otel.Tracer("server.controller")
