package order

import (
	"context"
	rpc "github.com/materials-resources/microservices-proto/golang/order"
	"github.com/materials-resources/s-prophet/app"
	"github.com/materials-resources/s-prophet/internal/order/api"
	"github.com/materials-resources/s-prophet/internal/order/data"
	"github.com/materials-resources/s-prophet/internal/order/service"
)

func init() {
	app.OnStart(
		"order.start", func(ctx context.Context, a *app.App) error {
			rpc.RegisterOrderServiceServer(
				a.GetGrpcServer(),
				api.NewOrderApi(*service.NewOrderController(data.NewModels(a.GetDB()))))
			return nil
		},
	)
}
