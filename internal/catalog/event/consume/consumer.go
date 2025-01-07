package consume

import (
	"github.com/materials-resources/s-prophet/app"
	"github.com/materials-resources/s-prophet/internal/catalog/event"
	"github.com/materials-resources/s-prophet/internal/catalog/repository"
	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/twmb/franz-go/pkg/sr"
)

type Consumer struct {
	serde  *sr.Serde
	client *kgo.Client

	cg *ConsumerGroup
}

func NewConsumer(manager *event.Manager, a *app.App, repo *repository.Repository) *Consumer {

	workers := NewWorkers(manager.Serde, manager.GetKotelTracer(), repository.NewRepository(a.GetDB()))
	cg := NewConsumerGroup()

	cg.AddWorker(event.DeleteProductTopic, workers.DeleteProduct)
	cg.AddWorker(event.UpdateGroupTopic, workers.UpdateGroup)
	cg.AddWorker(event.UpdateProductTopic, workers.UpdateProduct)

	opts := []kgo.Opt{
		kgo.ConsumerGroup("s-catalog-consumer-group"),
		kgo.ConsumeTopics(event.DeleteProductTopic, event.UpdateGroupTopic, event.UpdateProductTopic),
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
