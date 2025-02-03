package kafka

import (
	"fmt"
	"github.com/materials-resources/s-prophet/config"
	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/twmb/franz-go/plugin/kotel"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

type Factory struct {
	Config      config.Kafka
	kotelTracer *kotel.Tracer
	kotelMeter  *kotel.Meter
}

func NewFactory(config config.Kafka, tp trace.TracerProvider, mp metric.MeterProvider) *Factory {
	return &Factory{
		Config:      config,
		kotelTracer: newKotelTracer(tp),
		kotelMeter:  newKotelMeter(mp),
	}
}

// NewKafkaClient creates a new Kafka client with the same stored
// Kafka configuration.
func (s *Factory) NewKafkaClient(
	clientID string,
	additionalOpts ...kgo.Opt,
) (*kgo.Client, error) {
	// Create a new kotel service.
	kotelService := newKotel(s.kotelTracer, s.kotelMeter)

	kgoOpts, err := NewKgoConfig(&s.Config, kotelService)
	if err != nil {
		return nil, fmt.Errorf("failed to create a valid kafka client config: %w", err)
	}

	kgoOpts = append(kgoOpts, kgo.ClientID(clientID))
	kgoOpts = append(kgoOpts, additionalOpts...)

	kafkaClient, err := kgo.NewClient(kgoOpts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create kafka client: %w", err)
	}

	return kafkaClient, nil
}

func (s *Factory) GetKotelTracer() *kotel.Tracer {
	return s.kotelTracer
}

func newKotelTracer(tracerProvider trace.TracerProvider) *kotel.Tracer {
	// Create a new kotel tracer with the provided tracer provider and
	// propagator.
	tracerOpts := []kotel.TracerOpt{
		kotel.TracerProvider(tracerProvider),
		kotel.TracerPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{})),
	}
	return kotel.NewTracer(tracerOpts...)
}

func newKotelMeter(meterProvider metric.MeterProvider) *kotel.Meter {
	// Create a new kotel meter using a provided meter provider.
	meterOpts := []kotel.MeterOpt{kotel.MeterProvider(meterProvider)}
	return kotel.NewMeter(meterOpts...)
}

func newKotel(tracer *kotel.Tracer, meter *kotel.Meter) *kotel.Kotel {
	kotelOps := []kotel.Opt{
		kotel.WithTracer(tracer),
		kotel.WithMeter(meter),
	}
	return kotel.NewKotel(kotelOps...)
}
