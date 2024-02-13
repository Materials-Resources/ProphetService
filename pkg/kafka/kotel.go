package kafka

import (
	"github.com/twmb/franz-go/plugin/kotel"
	"go.opentelemetry.io/otel/propagation"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

func NewKotelTracer(tp *tracesdk.TracerProvider) *kotel.Tracer {

	traceOpts := []kotel.TracerOpt{
		kotel.TracerProvider(tp),
		kotel.TracerPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{})),
	}
	return kotel.NewTracer(traceOpts...)
}

func NewKotel(tracer *kotel.Tracer) *kotel.Kotel {
	kotelOpts := []kotel.Opt{
		kotel.WithTracer(tracer),
	}
	return kotel.NewKotel(kotelOpts...)
}
