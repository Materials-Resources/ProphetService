package customer

import (
	"context"
	"github.com/materials-resources/s-prophet/app"
	"github.com/materials-resources/s-prophet/internal/customer/api"
	"github.com/materials-resources/s-prophet/internal/customer/repository"
	"github.com/materials-resources/s-prophet/internal/customer/service"
	"github.com/materials-resources/s-prophet/internal/grpc/customer/customerconnect"
)

func init() {
	app.OnStart(
		"customer.start", func(ctx context.Context, a *app.App) error {
			svc := service.NewCustomerService(repository.NewRepository(a.GetDB()))

			path, handler := customerconnect.NewCustomerServiceHandler(api.NewCustomerServiceHandler(svc))

			a.RegisterService(path, handler, customerconnect.CustomerServiceName)
			return nil
		},
	)
}
