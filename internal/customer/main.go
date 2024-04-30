package customer

import (
	"context"
	"github.com/materials-resources/s_prophet/app"
	"github.com/materials-resources/s_prophet/infrastructure/data"
	"github.com/materials-resources/s_prophet/internal/customer/api"
	"github.com/materials-resources/s_prophet/internal/customer/service"
	svc "github.com/materials-resources/s_prophet/proto/customer/v1"
)

func init() {
	app.OnStart(
		"CustomerService.init", func(ctx context.Context, a *app.App) error {
			m := data.NewModels(a.GetDB())
			svc.RegisterCustomerServiceServer(
				a.GetServer(),
				api.NewCustomerServer(service.NewCustomerService(m)),
			)
			return nil
		},
	)
}
