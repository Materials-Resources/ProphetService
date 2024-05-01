package catalog

import (
	"context"
	"fmt"
	"github.com/materials-resources/s-prophet/app"
	"github.com/materials-resources/s-prophet/internal/catalog/api"
	"github.com/materials-resources/s-prophet/internal/catalog/service"
	"github.com/materials-resources/s-prophet/pkg/kafka"
	svc "github.com/materials-resources/s-prophet/proto/catalog/v1"
	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/twmb/franz-go/pkg/sr"
)

func init() {
	app.OnStart(
		"catalogService.init", func(ctx context.Context, a *app.App) error {

			var serde sr.Serde
			service.RegisterSchema(&serde)

			kotelTracer := kafka.NewKotelTracer(a.GetTP())

			kotel := kafka.NewKotel(kotelTracer)

			client, err := kgo.NewClient(
				kgo.SeedBrokers(a.Config.Broker.Host), kgo.WithHooks(kotel.Hooks()...),
				kgo.ConsumeTopics(service.UpdateProductTopic),
			)

			if err != nil {
				fmt.Println(err)
			}

			cs := service.NewCatalogService(a.GetModels(), a.GetTP().Tracer("CatalogApi"), client, &serde)

			cs.RegisterWorkers()

			svc.RegisterCatalogServiceServer(
				a.GetGrpcServer(), api.NewCatalogApi(
					cs, a.GetTP().Tracer("CatalogApi"), &service.KafkaProducer{Client: client, Serde: &serde}),
			)

			return nil
		},
	)
}
