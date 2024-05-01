package app

import (
	"context"
	"github.com/materials-resources/s-prophet/infrastructure/data"
	"github.com/rs/zerolog"
	"net"
	"sync"

	"github.com/uptrace/bun"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/metric"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	"google.golang.org/grpc"
)

type App struct {
	Config *Config

	server     *grpc.Server
	serverOnce sync.Once

	tp *tracesdk.TracerProvider
	mp *metric.MeterProvider

	bunOnce sync.Once
	bun     *bun.DB

	log *zerolog.Logger
}

func NewApp(config *Config) (*App, error) {

	a := &App{
		Config: config,
	}

	a.tp = a.newTracer()

	a.newLogger()

	otel.SetTracerProvider(a.tp)

	return a, nil
}

func (a *App) Start() {
	ctx := context.Background()

	if err := onStart.Run(ctx, a); err != nil {
		a.Logger().Fatal().Err(err).Msg("failed to run onStart hooks")
	}

	lis, err := net.Listen(
		"tcp",
		"0.0.0.0:50058",
	)
	if err != nil {
		a.Logger().Fatal().Err(err).Msg("failed to create listener")

	}

	err = a.GetGrpcServer().Serve(lis)
	if err != nil {
		a.Logger().Fatal().Err(err).Msg("failed to start server")
	}

}

// Stop stops the application.
func (a *App) Stop() {
	a.GetGrpcServer().Stop()
}

func (a *App) GetGrpcServer() *grpc.Server {
	a.serverOnce.Do(
		func() {
			a.server = a.newGrpcServer()
		},
	)
	return a.server
}

// GetDB returns an initialized instance of  bun.DB.
func (a *App) GetDB() *bun.DB {

	a.bunOnce.Do(
		func() {
			a.newBun()
		})

	return a.bun
}

// GetTP returns initialized instance of tracesdk.TracerProvider.
func (a *App) GetTP() *tracesdk.TracerProvider {
	return a.tp
}

// GetModels returns an initialized instance of data.Models.
func (a *App) GetModels() data.Models {

	return *data.NewModels(a.GetDB())
}
func (a *App) GetMP() *metric.MeterProvider {
	return a.mp
}
