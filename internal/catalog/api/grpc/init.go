package catalog

import (
	"context"
	"fmt"
	"github.com/materials-resources/s_prophet/app"
	"github.com/materials-resources/s_prophet/infrastructure/db/prophet_21_1_4559"
	"github.com/materials-resources/s_prophet/internal/catalog/service"
	"github.com/materials-resources/s_prophet/pkg/kafka"
	svc "github.com/materials-resources/s_prophet/proto/catalog/v1alpha0"
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
				kgo.SeedBrokers(a.Config.App.Events.Brokers...), kgo.WithHooks(kotel.Hooks()...),
				kgo.ConsumeTopics(service.UpdateProductTopic),
			)

			m := prophet_21_1_4559.NewModels(a.GetDB())

			if err != nil {
				fmt.Println(err)
			}

			cs := service.NewCatalogService(*m, a.GetTP().Tracer("CatalogService"), client, &serde)

			cs.RegisterWorkers()

			svc.RegisterCatalogServiceServer(
				a.GetServer(), &catalogService{
					producer: &service.KafkaProducer{Client: client, Serde: &serde},
					tracer:   a.GetTP().Tracer("CatalogService"),
					service:  cs,
				},
			)

			return nil
		},
	)
}
