package inventory

import (
	"context"

	"github.com/materials-resources/s_prophet/app"
	"github.com/materials-resources/s_prophet/internal/inventory/repository"
	svc "github.com/materials-resources/s_prophet/proto/inventory/v1alpha0"
)

func init() {
	app.OnStart("InventoryService.init", func(ctx context.Context, a *app.App) error {
		svc.RegisterInventoryServiceServer(a.GetServer(),
			&inventoryService{repo: repository.NewBunInventoryRepository(a.GetDB())},
		)
		return nil
	},
	)
}
