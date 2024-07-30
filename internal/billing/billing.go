package billing

import (
	"github.com/materials-resources/s-prophet/app"
	"github.com/materials-resources/s-prophet/internal/billing/adapters/grpc"
	"golang.org/x/net/context"
)

func init() {
	app.OnStart(
		"catalogService.start", func(ctx context.Context, a *app.App) error {
			grpc.NewServerAdapter(a)

			return nil
		},
	)
}
