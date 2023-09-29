package server

import (
	"github.com/materials-resources/s_prophet/pkg/internal/core/server"
	"github.com/materials-resources/s_prophet/pkg/internal/core/service/order"
	"github.com/materials-resources/s_prophet/pkg/internal/core/service/product"
	"github.com/materials-resources/s_prophet/pkg/internal/core/service/receiving"
	"github.com/materials-resources/s_prophet/pkg/internal/core/service/shipping"
	"github.com/materials-resources/s_prophet/pkg/internal/infra/repository"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Serve() {
	s := grpc.NewServer()

	db, _ := repository.NewDB()

	//Define repositories
	orderRepo := repository.NewOrderRepository(db)
	productRepo := repository.NewProductRepository(db)
	shippingRepo := repository.NewShippingRepository(db)
	receivingRepo := repository.NewReceivingRepository(db)

	//Register the grpc servers
	order.NewOrderServer(
		s,
		orderRepo,
	)
	product.NewProductServer(
		s,
		productRepo,
	)
	shipping.NewShippingServer(
		s,
		shippingRepo,
	)

	receiving.NewReceivingServer(
		s,
		receivingRepo,
	)

	//Enable GRPC Reflection for clients
	reflection.Register(s)

	grpcServer := server.NewGrpcServer(s)

	grpcServer.Start()
}
