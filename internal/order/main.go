package order

import (
	"context"
	"github.com/materials-resources/s_prophet/app"
	"github.com/materials-resources/s_prophet/infrastructure/data"
	"github.com/materials-resources/s_prophet/internal/order/api"
	"github.com/materials-resources/s_prophet/internal/order/service"
	svc "github.com/materials-resources/s_prophet/proto/order/v1alpha0"
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
