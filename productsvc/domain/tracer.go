package domain

import "go.opentelemetry.io/otel"

var Tracer = otel.Tracer("product-svc")

