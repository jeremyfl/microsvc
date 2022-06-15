package domain

import "go.opentelemetry.io/otel"

var Tracer = otel.Tracer("stocksvc-listener")
