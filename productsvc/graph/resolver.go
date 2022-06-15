package graph

import (
	"gitlab.com/jeremylo/microsvc/productsvc/domain"
	"go.opentelemetry.io/otel/trace"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Services domain.Services
	Tracer trace.Tracer
}
