package cmd

import (
	"github.com/materials-resources/s_prophet/controller"
	"github.com/materials-resources/s_prophet/core/server"
	"github.com/materials-resources/s_prophet/repository"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var serveCmd = &cobra.Command{
	Use: "serve",
	Run: func(cmd *cobra.Command, args []string) {
		s := grpc.NewServer()

		db, _ := repository.NewDB()

		//Define repositories
		orderRepo := repository.NewOrderRepository(db)
		productRepo := repository.NewProductRepository(db)
		shippingRepo := repository.NewShippingRepository(db)
		receivingRepo := repository.NewReceivingRepository(db)

		//Register the grpc servers
		controller.NewOrderService(
			s,
			orderRepo,
		)
		controller.NewProductService(
			s,
			productRepo,
		)
		controller.NewShippingService(
			s,
			shippingRepo,
		)

		controller.NewReceivingService(
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
