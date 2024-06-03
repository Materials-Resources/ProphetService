package grpc

import (
	"github.com/materials-resources/s-prophet/app"
	"github.com/materials-resources/s-prophet/internal/billing/core/api"
	"google.golang.org/grpc"
)

type ServerAdapter struct {
	server *grpc.Server
}

func NewServerAdapter(app *app.App) *ServerAdapter {

	a := api.NewBillingAdapter(app.GetModels())

	billingS := BillingService{
		api: a,
	}
	billingS.registerService(app.GetGrpcServer())
	app.Logger().Info().Msg("Billing Service registered")

	return &ServerAdapter{
		server: app.GetGrpcServer(),
	}
}
