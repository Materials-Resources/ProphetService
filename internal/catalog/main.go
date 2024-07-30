package catalog

import (
	"context"
	rpc "github.com/materials-resources/microservices-proto/golang/catalog"
	"github.com/materials-resources/s-prophet/app"
	"github.com/materials-resources/s-prophet/internal/catalog/api"
	"github.com/materials-resources/s-prophet/internal/catalog/core/event"
	"github.com/materials-resources/s-prophet/internal/catalog/core/event/consume"
	"github.com/materials-resources/s-prophet/internal/catalog/core/event/produce"
	"github.com/materials-resources/s-prophet/internal/catalog/core/service"
)

func init() {
	app.OnStart(
		"catalog.start", func(ctx context.Context, a *app.App) error {

			manager, err := event.NewManager(a)
			if err != nil {
				return err
			}

			err = manager.RegisterSchemasWithSerde()
			if err != nil {
				return err
			}

			producer := produce.NewProducer(a, manager)

			cs := service.NewCatalogService(a, producer)

			consumer := consume.NewConsumer(manager, a, cs)

			rpc.RegisterCatalogServiceServer(
				a.GetGrpcServer(), api.NewCatalogApi(cs,
					a.GetTP().Tracer("CatalogApi"),
				))

			go func() { consumer.Consume() }()

			return nil
		},
	)

	app.OnSetup("catalog.setup", func(ctx context.Context, a *app.App) error {
		manager, err := event.NewManager(a)
		if err != nil {
			return err
		}
		err = manager.RegisterSchemasWithRegistry()
		if err != nil {
			return err
		}

		err = manager.RegisterTopics()
		if err != nil {
			return err
		}

		return nil
	})
}
