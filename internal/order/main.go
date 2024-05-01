package order

import (
	"context"
	"github.com/materials-resources/s-prophet/app"
	"github.com/materials-resources/s-prophet/infrastructure/data"
	"github.com/materials-resources/s-prophet/internal/order/api"
	"github.com/materials-resources/s-prophet/internal/order/service"
	svc "github.com/materials-resources/s-prophet/proto/order/v1"
)

func init() {
	app.OnStart(
		"OrderService.init", func(ctx context.Context, a *app.App) error {
			m := data.NewModels(a.GetDB())
			svc.RegisterOrderServiceServer(
				a.GetServer(),
				api.NewOrderApi(service.NewOrderService(*m)),
			)
			return nil
		},
	)
}
