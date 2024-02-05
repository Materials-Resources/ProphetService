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
	"github.com/uptrace/bun/extra/bunotel"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	service     = "prophet"
	environment = "development"
	id          = 1
)

var serveCmd = &cobra.Command{
	Use: "serve",
	Run: func(cmd *cobra.Command, args []string) {
		cPath := rootCmd.PersistentFlags().Lookup("config").Value

		app.NewConfig(cPath.String())
		s := grpc.NewServer(grpc.StatsHandler(otelgrpc.NewServerHandler()))

		tp, err := tracerProvider("http://localhost:14268/api/traces")

		otel.SetTracerProvider(tp)

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

		bundb.AddQueryHook(bunotel.NewQueryHook(bunotel.WithDBName("mydb"))) //bundb.AddQueryHook(bundebug.
		// NewQueryHook(bundebug.WithVerbose(true)))

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

func tracerProvider(url string) (*tracesdk.TracerProvider, error) {
	// Create the Jaeger exporter
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return nil, err
	}
	tp := tracesdk.NewTracerProvider(
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exp),
		// Record information about this application in a Resource.
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(service),
			attribute.String("environment", environment),
			attribute.Int64("ID", id),
		),
		),
	)
	return tp, nil
}
