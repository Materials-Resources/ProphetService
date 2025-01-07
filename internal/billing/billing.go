package billing

import (
	"github.com/materials-resources/s-prophet/app"
	"github.com/materials-resources/s-prophet/internal/billing/api"
	"github.com/materials-resources/s-prophet/internal/billing/data"
	"github.com/materials-resources/s-prophet/internal/billing/service"
	"golang.org/x/net/context"
)

func init() {
	app.OnStart(
		"billing", func(ctx context.Context, a *app.App) error {

			api.RegisterRoutes(
				a.GetGrpcServer(),
				api.NewBillingRoutes(service.NewBilling(data.NewModels(a.GetDB()))))

			return nil
		},
	)
}
