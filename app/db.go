package app

import (
	"database/sql"
	"github.com/materials-resources/s-prophet/config"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	"net/url"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mssqldialect"
	"github.com/uptrace/bun/extra/bundebug"
	"github.com/uptrace/bun/extra/bunotel"
)

func (a *App) newBun() {
	query := url.Values{}
	query.Add("database", a.conf.Database.Name)
	u := &url.URL{
		Scheme:   "sqlserver",
		User:     url.UserPassword(a.conf.Database.User, a.conf.Database.Password),
		Host:     a.conf.Database.Host,
		RawQuery: query.Encode(),
	}

	db, err := sql.Open(
		"sqlserver",
		u.String(),
	)

	if err != nil {
		a.Logger.Fatal().Err(err).Msg("could not connect to database")
	}

	if err := db.Ping(); err != nil {
		a.Logger.Fatal().Err(err).Msg("could not ping database")
	}

	bunDb := bun.NewDB(
		db,
		mssqldialect.New(),
	)

	a.Logger.Info().Str("db", query.Get("database")).Str("host", u.Hostname()).Msg("connected to database")

	registerBunOtelTracer(bunDb, a.Otel.GetTracerProvider())

	if a.conf.Environment == config.EnvironmentDevelopment {
		a.Logger.Info().Msg("registering bun debug hooks")
		bunDb.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
	}

	a.bun = bunDb
}

func registerBunOtelTracer(bunDb *bun.DB, tp *tracesdk.TracerProvider) {
	bunDb.AddQueryHook(
		bunotel.NewQueryHook(
			bunotel.WithTracerProvider(tp),
		),
	)
}
