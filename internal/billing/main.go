package billing

import (
	"github.com/materials-resources/s-prophet/app"
	"github.com/materials-resources/s-prophet/internal/billing/api"
	"github.com/materials-resources/s-prophet/internal/billing/service"
	"github.com/materials-resources/s-prophet/internal/grpc/billing/billingconnect"
	"golang.org/x/net/context"
)

func init() {
	app.OnStart(
		"billing.start", func(ctx context.Context, a *app.App) error {

			svc := service.NewBillingService(a)

			path, handler := billingconnect.NewBillingServiceHandler(api.NewBillingServiceHandler(svc))

			a.RegisterService(path, handler, billingconnect.BillingServiceName)

			return nil
		},
	)
}
