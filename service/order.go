package service

import (
	"context"

	"github.com/materials-resources/s_prophet/core/port/repository"
	rpc "github.com/materials-resources/s_prophet/proto/order/v1alpha0"
	"google.golang.org/grpc"
)

type orderService struct {
	orderRepository repository.OrderRepository
}

func NewOrderServer(server *grpc.Server, orderRepository repository.OrderRepository) {
	orderServer := &orderService{orderRepository: orderRepository}
	rpc.RegisterOrderServiceServer(server, orderServer)
}

func (s *orderService) GetOrder(ctx context.Context, req *rpc.GetOrderRequest) (*rpc.GetOrderResponse, error) {
	return s.orderRepository.Read(ctx, req.GetId())
}
