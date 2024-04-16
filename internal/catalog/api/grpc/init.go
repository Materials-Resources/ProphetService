package catalog

import (
	"context"
	"github.com/materials-resources/s_prophet/infrastructure/db/prophet_21_1_4559"
	"github.com/materials-resources/s_prophet/internal/catalog/service"
	"github.com/twmb/franz-go/pkg/sr"

	"github.com/materials-resources/s_prophet/app"
	kafka2 "github.com/materials-resources/s_prophet/internal/catalog/kafka"
	"github.com/materials-resources/s_prophet/internal/catalog/repository"
	"github.com/materials-resources/s_prophet/pkg/kafka"
	svc "github.com/materials-resources/s_prophet/proto/catalog/v1alpha0"
)

func init() {
	app.OnStart(
		"catalogService.init", func(ctx context.Context, a *app.App) error {
			topic := "prophet.catalog"

			kotelTracer := kafka.NewKotelTracer(a.GetTP())

			kotel := kafka.NewKotel(kotelTracer)

			repo := repository.NewBunCatalogRepository(a.GetDB(), a.GetTP())

			productWorker := kafka2.NewProductWorker(repo)
			var serde sr.Serde
			service.RegisterSchema(&serde)

			cg := kafka.NewConsumerGroup(productWorker.ConsumeDeleteProduct, kotelTracer)
			cg.Start(a.Config.App.Events.Brokers, topic, "14", kotel)

			producer := kafka2.NewProducer(a.Config.App.Events.Brokers, topic, kotel, a.GetTP().Tracer("producer"))
			// go consumer1.DeleteProduct()

			m := prophet_21_1_4559.NewModels(a.GetDB())
			svc.RegisterCatalogServiceServer(
				a.GetServer(), &catalogService{
					producer: *producer,
					tracer:   a.GetTP().Tracer("CatalogService"),
					service:  service.NewCatalogService(*m),
				},
			)
			return nil
		},
	)
}
