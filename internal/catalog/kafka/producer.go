package kafka

import (
	"context"
	"fmt"
	"sync"

	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/twmb/franz-go/plugin/kotel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

type Producer struct {
	client *kgo.Client
	topic  string
	tracer trace.Tracer
}

func NewProducer(brokers []string, topic string, kotel *kotel.Kotel, tracer trace.Tracer) *Producer {

	client, err := kgo.NewClient(
		kgo.SeedBrokers(brokers...), kgo.WithHooks(kotel.Hooks()...),
	)
	if err != nil {
		panic(err)
	}
	return &Producer{client: client, topic: topic, tracer: tracer}
}

func (p *Producer) PublishDelete(ctx context.Context, id string) {
	_, span := p.tracer.Start(ctx, "Publish Delete")
	defer span.End()

	var wg sync.WaitGroup
	wg.Add(1)

	p.client.Produce(ctx, &kgo.Record{Topic: p.topic, Value: []byte(id)}, func(_ *kgo.Record, err error) {
		defer wg.Done()
		if err != nil {
			fmt.Printf("record had a produce error: %v\n", err)
			// Set the status and record error on the span.
			span.SetStatus(codes.Error, err.Error())
			span.RecordError(err)
		}
	},
	)
	wg.Wait()
}

func (p *Producer) Close() {
	p.client.Close()
}
