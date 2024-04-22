package inventory

import (
	"context"
	"github.com/materials-resources/s_prophet/app"
	"github.com/materials-resources/s_prophet/infrastructure/data"
	"github.com/materials-resources/s_prophet/internal/inventory/api"
	"github.com/materials-resources/s_prophet/internal/inventory/service"
	svc "github.com/materials-resources/s_prophet/proto/inventory/v1alpha0"
)

func init() {
	app.OnStart(
		"InventoryService.init", func(ctx context.Context, a *app.App) error {

			m := data.NewModels(a.GetDB())
			svc.RegisterInventoryServiceServer(
				a.GetServer(),
				api.NewInventoryApi(service.NewInventoryService(*m)),
			)
			return nil
		},
	)
}
