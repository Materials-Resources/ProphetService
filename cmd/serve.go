package cmd

import (
	"github.com/materials-resources/s_prophet/config"
	"github.com/materials-resources/s_prophet/core/server"
	"github.com/materials-resources/s_prophet/repository"
	"github.com/materials-resources/s_prophet/service"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var serveCmd = &cobra.Command{
	Use: "serve",
	Run: func(cmd *cobra.Command, args []string) {
		cPath := rootCmd.PersistentFlags().Lookup("config").Value

		config.NewConfig(cPath.String())
		s := grpc.NewServer()

		db, _ := repository.NewDB(&config.Configuration)

		//Define repositories
		orderRepo := repository.NewOrderRepository(db)
		productRepo := repository.NewProductRepository(db)
		shippingRepo := repository.NewShippingRepository(db)
		receivingRepo := repository.NewReceivingRepository(db)

		//Register the grpc servers
		service.NewOrderServer(
			s,
			orderRepo,
		)
		service.NewProductServer(
			s,
			productRepo,
		)
		service.NewShippingService(
			s,
			shippingRepo,
		)

		service.NewReceivingService(
			s,
			receivingRepo,
		)

		//Enable GRPC Reflection for clients
		reflection.Register(s)

		grpcServer := server.NewGrpcServer(s)

		grpcServer.Start()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
