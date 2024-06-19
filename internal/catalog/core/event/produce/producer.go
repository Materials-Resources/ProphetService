package produce

import (
	"context"
	"fmt"
	"github.com/materials-resources/s-prophet/app"
	"github.com/materials-resources/s-prophet/internal/catalog/core/event"
	"github.com/materials-resources/s-prophet/internal/catalog/domain"
	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/twmb/franz-go/pkg/sr"
	"go.opentelemetry.io/otel/trace"
	"sync"
)

type Producer struct {
	tracer trace.Tracer
	serde  *sr.Serde
	client *kgo.Client
}

func NewProducer(app *app.App, manager *event.Manager) *Producer {
	client, err := kgo.NewClient(manager.GetDefaultKgoOptions()...)
	if err != nil {
		panic(err)

	}
	return &Producer{client: client, serde: manager.Serde, tracer: app.GetTP().Tracer("Producer")}
}

func (p *Producer) DeleteProduct(ctx context.Context, product *domain.Product) error {

	var wg sync.WaitGroup
	wg.Add(1)

	rec := &kgo.Record{
		Topic: event.DeleteProductTopic, Value: p.serde.MustEncode(event.ProductRecordFromDomain(product)),
	}

	p.client.Produce(ctx, rec, func(record *kgo.Record, err error) {
		defer wg.Done()
		if err != nil {
			fmt.Printf("record had a produce error: %v\n", err)
		}

	})
	wg.Wait()
	return nil
}

func (p *Producer) UpdateGroup(ctx context.Context, group *domain.ProductGroup) error {

	var wg sync.WaitGroup
	wg.Add(1)

	rec := &kgo.Record{
		Topic: event.UpdateGroupTopic, Value: p.serde.MustEncode(event.ProductGroupRecordFromDomain(group)),
	}

	p.client.Produce(ctx, rec, func(record *kgo.Record, err error) {
		defer wg.Done()
		if err != nil {
			fmt.Printf("record had a produce error: %v\n", err)
		}

	})
	wg.Wait()
	return nil
}
