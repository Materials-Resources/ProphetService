package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
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

	db     *bun.DB
	dbOnce sync.Once

	server     *grpc.Server
	serverOnce sync.Once

	tp *tracesdk.TracerProvider
	mp *metric.MeterProvider
}

func NewApp(config *Config) (*App, error) {

	tp, err := newTraceProvider(config.Tracing.Service, config.Tracing.Environment)
	if err != nil {
		return nil, errors.New("could not create tracing provider")
	}

	otel.SetTracerProvider(tp)

	return &App{
		Config: config,
		tp:     tp,
	}, nil
}

func (a *App) Start() {
	ctx := context.Background()

	if err := onStart.Run(ctx, a); err != nil {
		log.Err(err)
		return
	}

	listener, err := net.Listen(
		"tcp",
		"0.0.0.0:50058",
	)
	if err != nil {
		log.Err(err).Msg("failed to create listener")
		return

	}

	err = a.server.Serve(listener)
	if err != nil {
		fmt.Println(err)
		return
	}

}

func (a *App) Stop() {

}

func (a *App) GetServer() *grpc.Server {
	a.serverOnce.Do(
		func() {
			a.server = newGrpcServer()
		},
	)
	return a.server
}
func (a *App) GetDB() *bun.DB {

	a.dbOnce.Do(
		func() {
			a.db = a.newBunDB()
		},
	)

	return a.db
}

func (a *App) GetTP() *tracesdk.TracerProvider {
	return a.tp
}

func (a *App) GetMP() *metric.MeterProvider {
	return a.mp
}
