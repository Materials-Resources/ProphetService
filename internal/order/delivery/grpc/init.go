package order

import (
	"github.com/materials-resources/s_prophet/app"
	"github.com/materials-resources/s_prophet/internal/order/repository"
	svc "github.com/materials-resources/s_prophet/proto/order/v1alpha0"
	"golang.org/x/net/context"
)

func init() {
	app.OnStart("OrderService.init", func(ctx context.Context, a *app.App) error {
		svc.RegisterOrderServiceServer(a.GetServer(),
			&orderService{repo: repository.NewBunOrderRepository(a.GetDB())},
		)
		return nil
	},
	)
}
