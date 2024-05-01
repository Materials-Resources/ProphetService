package order

import (
	"context"
	"github.com/materials-resources/s-prophet/app"
	"github.com/materials-resources/s-prophet/internal/order/api"
	"github.com/materials-resources/s-prophet/internal/order/service"
	svc "github.com/materials-resources/s-prophet/proto/order/v1"
)

func init() {
	app.OnStart(
		"OrderService.init", func(ctx context.Context, a *app.App) error {
			svc.RegisterOrderServiceServer(
				a.GetGrpcServer(),
				api.NewOrderApi(service.NewOrderService(a.GetModels())),
			)
			return nil
		},
	)
}
