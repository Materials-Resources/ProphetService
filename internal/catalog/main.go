package catalog

import (
	"context"
	"github.com/materials-resources/s-prophet/app"
	"github.com/materials-resources/s-prophet/internal/catalog/api"
	"github.com/materials-resources/s-prophet/internal/catalog/event"
	"github.com/materials-resources/s-prophet/internal/catalog/event/consume"
	"github.com/materials-resources/s-prophet/internal/catalog/event/produce"
	"github.com/materials-resources/s-prophet/internal/catalog/repository"
	"github.com/materials-resources/s-prophet/internal/catalog/service"
	"github.com/materials-resources/s-prophet/internal/grpc/catalog/catalogconnect"
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

			svc := service.NewCatalogService(a, producer)

			consumer := consume.NewConsumer(manager, a, repository.NewRepository(a.GetDB()))

			path, handler := catalogconnect.NewCatalogServiceHandler(api.NewCatalogServiceHandler(svc))

			a.RegisterService(path, handler, catalogconnect.CatalogServiceName)

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
