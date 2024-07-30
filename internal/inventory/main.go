package inventory

import (
	"context"
	"github.com/materials-resources/s-prophet/app"
	"github.com/materials-resources/s-prophet/internal/inventory/api"
	"github.com/materials-resources/s-prophet/internal/inventory/service"
	svc "github.com/materials-resources/s-prophet/proto/inventory/v1"
)

func init() {
	app.OnStart(
		"inventory.start", func(ctx context.Context, a *app.App) error {
			a.Logger.Info().Msg("Initializing Inventory Service")
			svc.RegisterInventoryServiceServer(
				a.GetGrpcServer(),
				api.NewInventoryApi(service.NewInventoryService(a.GetModels())),
			)
			return nil
		},
	)
}
