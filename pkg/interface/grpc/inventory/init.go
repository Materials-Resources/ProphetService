package inventory

import (
	"context"

	"github.com/materials-resources/s_prophet/app"
	"github.com/materials-resources/s_prophet/pkg/infrastructure/db/prophet_19_1_3668"
	svc "github.com/materials-resources/s_prophet/proto/inventory/v1alpha0"
)

func init() {
	app.OnStart("InventoryService.init", func(ctx context.Context, a *app.App) error {
		svc.RegisterInventoryServiceServer(a.GetServer(), &inventoryService{repo: prophet_19_1_3668.
			NewBunInventoryRepository(a.GetDB()),
		},
		)
		return nil
	},
	)
}
