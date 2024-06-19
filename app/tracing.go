package app

import (
	"context"
	"go.opentelemetry.io/otel/trace"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.24.0"
)

// newTracer creates a new tracer provider and sets it on the App.
func (a *App) newTracer() *tracesdk.TracerProvider {

	ctx := context.Background()

	exp, err := otlptracegrpc.New(ctx, otlptracegrpc.WithInsecure(), otlptracegrpc.WithEndpoint("localhost:4317"))

	if err != nil {
		a.Logger().Err(err).Msg("could not create tracing exporter")
	}

	tp := tracesdk.NewTracerProvider(
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exp),
		// Record information about this application in a Resource.
		tracesdk.WithResource(
			resource.NewWithAttributes(
				semconv.SchemaURL,
				semconv.ServiceNameKey.String(a.Config.Tracing.Service),
				attribute.String("environment", a.Config.App.Environment),
			),
		),
	)

	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(
			propagation.TraceContext{},
			propagation.Baggage{},
		),
	)

	return tp

}

// GetTP returns initialized instance of tracesdk.TracerProvider.
func (a *App) GetTP() *tracesdk.TracerProvider {
	return a.tp
}

func (a *App) CreateTracer(name string) trace.Tracer {
	return a.traceProvider.Tracer(name)
}
