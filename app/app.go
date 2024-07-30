package app

import (
	"context"
	"fmt"
	"github.com/materials-resources/s-prophet/config"
	"github.com/materials-resources/s-prophet/pkg/models"
	"github.com/rs/zerolog"
	"net"
	"os"
	"sync"

	"github.com/uptrace/bun"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/metric"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	"google.golang.org/grpc"
)

type App struct {
	conf *config.Config

	server     *grpc.Server
	serverOnce sync.Once

	traceProvider *tracesdk.TracerProvider
	mp            *metric.MeterProvider

	bunOnce sync.Once
	bun     *bun.DB

	Logger *zerolog.Logger
}

func NewApp(config *config.Config) (*App, error) {

	a := &App{
		conf: config,
	}

	a.initLog()

	err := a.newTraceProvider()
	if err != nil {
		return nil, err
	}
	otel.SetTracerProvider(a.traceProvider)

	return a, nil
}

// Setup runs the setup hooks.
func (a *App) Setup() error {
	ctx := context.Background()

	if err := onSetup.Run(ctx, a); err != nil {
		a.Logger.Fatal().Err(err).Msg("failed to run onSetup hooks")
	}

	return nil
}

func (a *App) Start() error {
	ctx := context.Background()

	if err := onStart.Run(ctx, a); err != nil {
		a.Logger.Fatal().Err(err).Msg("failed to run onStart hooks")
	}

	lis, err := net.Listen(
		"tcp",
		"0.0.0.0:50058",
	)

	if err != nil {
		a.Logger.Fatal().Err(err).Msg("failed to create listener")

	}

	err = a.GetGrpcServer().Serve(lis)
	if err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil

}

// Stop stops the application.
func (a *App) Stop() {
	a.server.Stop()
}

func (a *App) GetGrpcServer() *grpc.Server {
	a.serverOnce.Do(
		func() {
			a.server = a.newGrpcServer()
		},
	)
	return a.server
}

func (a *App) Config() *config.Config {
	return a.conf
}

// GetDB returns an initialized instance of  bun.DB.
func (a *App) GetDB() *bun.DB {

	a.bunOnce.Do(
		func() {
			a.newBun()
		})

	return a.bun
}

// GetModels returns an initialized instance of data.Models.
func (a *App) GetModels() models.Models {

	return *models.NewModels(a.GetDB())
}
func (a *App) GetMP() *metric.MeterProvider {
	return a.mp
}

func (a *App) initLog() {

	output := zerolog.ConsoleWriter{Out: os.Stderr}

	lg := zerolog.New(output).With().Timestamp().Logger()

	a.Logger = &lg
}
