package app

import (
	"database/sql"
	"fmt"
	"net/url"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mssqldialect"
	"github.com/uptrace/bun/extra/bunotel"
)

func (a *App) newBunDB() *bun.DB {
	query := url.Values{}
	query.Add("database", a.Config.Database.DB)
	u := &url.URL{Scheme: "sqlserver",
		User:     url.UserPassword(a.Config.Database.Username, a.Config.Database.Password),
		Host:     fmt.Sprintf("%s:%d", a.Config.Database.Host, a.Config.Database.Port),
		RawQuery: query.Encode(),
	}

	db, err := sql.Open(
		"sqlserver",
		u.String(),
	)

	if err != nil {

		fmt.Println(err)
		fmt.Println("could not connect to db")
	}

	bundb := bun.NewDB(
		db,
		mssqldialect.New(),
	)

	bundb.AddQueryHook(bunotel.NewQueryHook(bunotel.WithTracerProvider(a.GetTP()),
		bunotel.WithDBName(a.Config.Database.DB),
	),
	)

	return bundb
}
