package app

import (
	"context"
	"errors"
	"fmt"
	"net"
	"sync"

	"github.com/uptrace/bun"
	"go.opentelemetry.io/otel"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	"google.golang.org/grpc"
)

type App struct {
	config *Config

	db     *bun.DB
	dbOnce sync.Once

	server     *grpc.Server
	serverOnce sync.Once

	tp *tracesdk.TracerProvider
}

func NewApp(config *Config) (*App, error) {

	tp, err := newTraceProvider(config.Tracing.Service, config.Tracing.Environment)
	if err != nil {
		return nil, errors.New("could not create tracing provider")
	}

	otel.SetTracerProvider(tp)

	return &App{
		config: config,
		tp:     tp,
	}, nil
}

func (a *App) Start() {
	ctx := context.Background()

	if err := onStart.Run(ctx, a); err != nil {
		fmt.Errorf("%v", err)
		return
	}

	listener, err := net.Listen(
		"tcp",
		"0.0.0.0:50058",
	)

	err = a.server.Serve(listener)
	if err != nil {
		fmt.Println(err)
		return
	}

}

func (a *App) Stop() {

}

func (a *App) GetServer() *grpc.Server {
	a.serverOnce.Do(func() {
		a.server = newGrpcServer()
	},
	)
	return a.server
}
func (a *App) GetDB() *bun.DB {

	a.dbOnce.Do(func() {
		a.db = a.newBunDB()
	},
	)

	return a.db
}

func (a *App) GetTP() *tracesdk.TracerProvider {
	return a.tp
}
