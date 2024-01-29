package cmd

import (
	"database/sql"
	"fmt"
	"net/url"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/materials-resources/s_prophet/app"
	"github.com/materials-resources/s_prophet/core/server"
	"github.com/materials-resources/s_prophet/pkg/infrastructure/db/prophet_19_1_3668"
	grpc2 "github.com/materials-resources/s_prophet/pkg/interface/grpc"
	"github.com/spf13/cobra"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mssqldialect"
	"github.com/uptrace/bun/extra/bundebug"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var serveCmd = &cobra.Command{
	Use: "serve",
	Run: func(cmd *cobra.Command, args []string) {
		cPath := rootCmd.PersistentFlags().Lookup("config").Value

		app.NewConfig(cPath.String())
		s := grpc.NewServer()

		query := url.Values{}
		query.Add("database", app.Configuration.Database.DB)
		u := &url.URL{Scheme: "sqlserver",
			User:     url.UserPassword(app.Configuration.Database.Username, app.Configuration.Database.Password),
			Host:     fmt.Sprintf("%s:%d", app.Configuration.Database.Host, app.Configuration.Database.Port),
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
		bundb.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))

		catalogRepo := prophet_19_1_3668.NewBunCatalogRepository(bundb)
		orderRepo := prophet_19_1_3668.NewBunOrderRepository(bundb)
		inventoryRepo := prophet_19_1_3668.NewBunInventoryRepository(bundb)

		grpc2.NewCatalogService(s, catalogRepo)
		grpc2.NewOrderService(s, orderRepo)
		grpc2.NewInventoryService(s, inventoryRepo)

		//Enable GRPC Reflection for clients
		reflection.Register(s)

		grpcServer := server.NewGrpcServer(s)

		grpcServer.Start()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
