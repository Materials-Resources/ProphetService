package catalog

import (
	"context"
	rpc "github.com/materials-resources/microservices-proto/golang/catalog"
	"github.com/materials-resources/s-prophet/app"
	"github.com/materials-resources/s-prophet/internal/catalog/api"
	"github.com/materials-resources/s-prophet/internal/catalog/core/service"
)

func init() {
	app.OnStart(
		"catalogService.init", func(ctx context.Context, a *app.App) error {

			//var serde sr.Serde
			//service.RegisterSchema(&serde)
			//
			//kotelTracer := kafka.NewKotelTracer(a.GetTP())
			//
			//kotel := kafka.NewKotel(kotelTracer)
			//
			//client, err := kgo.NewClient(
			//	kgo.SeedBrokers(a.Config.Broker.Host), kgo.WithHooks(kotel.Hooks()...),
			//	kgo.ConsumeTopics(service.UpdateProductTopic),
			//)

			cs := service.NewCatalogService(a.GetModels(), a.GetTP().Tracer("CatalogApi"), a)

			//cs.RegisterWorkers()

			rpc.RegisterCatalogServiceServer(
				a.GetGrpcServer(), api.NewCatalogApi(cs,
					a.GetTP().Tracer("CatalogApi"),
				))

			return nil
		},
	)
}
