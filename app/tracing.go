package app

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.24.0"
)

type Tracer struct {
	tp *tracesdk.TracerProvider
}

func newTraceProvider(svcName, env string) (provider *tracesdk.TracerProvider, err error) {
	// Create the GRPC exporter
	ctx := context.Background()

	exp, err := otlptracegrpc.New(ctx)

	if err != nil {
		return nil, err
	}
	tp := tracesdk.NewTracerProvider(
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exp),
		// Record information about this application in a Resource.
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(svcName),
			attribute.String("environment", env),
		),
		),
	)
	return tp, nil
}
