package app

import (
	"context"
	"fmt"
	"github.com/materials-resources/s-prophet/config"
	"github.com/materials-resources/s-prophet/pkg/sr"
	"github.com/rs/zerolog"
	"net/http"
	"os"
	"sync"

	"github.com/uptrace/bun"
)

type App struct {
	conf *config.Config

	mux              *http.ServeMux
	reflectionRoutes []string

	bunOnce sync.Once
	bun     *bun.DB

	Otel          *Otel
	SchemaFactory *sr.Factory

	Logger *zerolog.Logger
}

func NewApp(config *config.Config) (*App, error) {

	mux := http.NewServeMux()

	otelProvider, err := NewOtel(*config)

	if err != nil {
		return nil, err
	}

	schemaFactory := sr.NewFactory(config.SchemaRegistry)

	a := &App{
		conf:          config,
		mux:           mux,
		Otel:          otelProvider,
		SchemaFactory: schemaFactory,
	}

	a.initLog()

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

	if a.conf.Environment.IsDevelopment() {
		a.registerReflection()
	}

	err := a.serveHttp()
	if err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil

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

func (a *App) initLog() {

	output := zerolog.ConsoleWriter{Out: os.Stderr}

	lg := zerolog.New(output).With().Timestamp().Logger()

	a.Logger = &lg
}
