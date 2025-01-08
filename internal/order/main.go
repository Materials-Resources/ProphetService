package order

import (
	"context"
	"github.com/materials-resources/s-prophet/app"
	"github.com/materials-resources/s-prophet/internal/grpc/order/orderconnect"
	"github.com/materials-resources/s-prophet/internal/order/api"
	"github.com/materials-resources/s-prophet/internal/order/repository"
	"github.com/materials-resources/s-prophet/internal/order/service"
)

func init() {
	app.OnStart(
		"order.start", func(ctx context.Context, a *app.App) error {
			path, handler := orderconnect.NewOrderServiceHandler(api.NewOrderServiceHandler(service.NewOrderService(repository.NewRepository(a.GetDB()))))
			a.RegisterService(path, handler, orderconnect.OrderServiceName)
			return nil
		},
	)
}
