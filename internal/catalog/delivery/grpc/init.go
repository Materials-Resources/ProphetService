package catalog

import (
	"context"

	"github.com/materials-resources/s_prophet/app"
	"github.com/materials-resources/s_prophet/internal/catalog/repository"

	svc "github.com/materials-resources/s_prophet/proto/catalog/v1alpha0"
)

func init() {
	app.OnStart("catalogService.init", func(ctx context.Context, a *app.App) error {
		svc.RegisterCatalogServiceServer(a.GetServer(), &catalogService{tracer: a.GetTP().Tracer("CatalogService"),
			repo: repository.NewBunCatalogRepository(a.GetDB(), a.GetTP()),
		},
		)
		return nil
	},
	)
}
