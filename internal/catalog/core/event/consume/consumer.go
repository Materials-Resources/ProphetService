package consume

import (
	"github.com/materials-resources/s-prophet/app"
	"github.com/materials-resources/s-prophet/internal/catalog/core/data"
	"github.com/materials-resources/s-prophet/internal/catalog/core/event"
	"github.com/materials-resources/s-prophet/internal/catalog/core/service"
	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/twmb/franz-go/pkg/sr"
)

type Consumer struct {
	serde  *sr.Serde
	client *kgo.Client

	cg *ConsumerGroup
}

func NewConsumer(manager *event.Manager, a *app.App, service service.CatalogService) *Consumer {

	workers := NewWorkers(manager.Serde, data.NewModels(a.GetDB()), manager.GetKotelTracer(), service)
	cg := NewConsumerGroup()

	cg.AddWorker(event.DeleteProductTopic, workers.DeleteProduct)
	cg.AddWorker(event.UpdateGroupTopic, workers.UpdateGroup)

	opts := []kgo.Opt{
		kgo.ConsumerGroup("s-catalog-consumer-group"),
		kgo.ConsumeTopics(event.DeleteProductTopic, event.UpdateGroupTopic),
		kgo.OnPartitionsAssigned(cg.assigned),
		kgo.OnPartitionsRevoked(cg.lost),
		kgo.OnPartitionsLost(cg.lost),
		kgo.DisableAutoCommit(),
		kgo.BlockRebalanceOnPoll(),
	}
	opts = append(opts, manager.GetDefaultKgoOptions()...)

	client, err := kgo.NewClient(opts...)
	if err != nil {
		panic(err)
	}

	return &Consumer{client: client, cg: cg}
}

func (c *Consumer) Consume() {
	c.cg.pollTopics(c.client)
}
