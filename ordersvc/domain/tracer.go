package domain

import "go.opentelemetry.io/otel"

var Tracer = otel.Tracer("ordersvc-api")
