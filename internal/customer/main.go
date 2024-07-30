package customer

import (
	"context"
	"github.com/materials-resources/s-prophet/app"
	"github.com/materials-resources/s-prophet/internal/customer/api"
	"github.com/materials-resources/s-prophet/internal/customer/service"
	"github.com/materials-resources/s-prophet/pkg/models"
	svc "github.com/materials-resources/s-prophet/proto/customer/v1"
)

func init() {
	app.OnStart(
		"customer.start", func(ctx context.Context, a *app.App) error {
			m := models.NewModels(a.GetDB())
			svc.RegisterCustomerServiceServer(
				a.GetGrpcServer(),
				api.NewCustomerServer(service.NewCustomerService(m)),
			)
			return nil
		},
	)
}
