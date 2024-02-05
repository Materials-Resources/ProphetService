package grpc

import (
	"context"

	"github.com/materials-resources/s_prophet/pkg/domain/entities"
	"github.com/materials-resources/s_prophet/pkg/domain/repositories"
	rpc "github.com/materials-resources/s_prophet/proto/order/v1alpha0"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewOrderService(server *grpc.Server, repo repositories.OrderRepository) {
	svc := &orderService{repo: repo}
	rpc.RegisterOrderServiceServer(server, svc)
}

type orderService struct {
	repo repositories.OrderRepository
}

func (s orderService) CreateOrder(ctx context.Context, request *rpc.CreateOrderRequest) (*rpc.CreateOrderResponse,
	error,
) {
	s.repo.CreateOrder(entities.ValidatedOrder{})
	return nil, nil
}

func (s orderService) GetOrder(ctx context.Context, request *rpc.GetOrderRequest) (*rpc.GetOrderResponse, error) {
	o, err := s.repo.FindOrderById(request.GetId())
	if err != nil {
		return nil, err
	}
	return ToPBGetOrderResponse(&o.Order)
}

func (s orderService) GetPickTicketById(
	ctx context.Context, request *rpc.GetPickTicketByIdRequest,
) (*rpc.GetPickTicketByIdResponse, error) {
	ent, err := s.repo.FindPickTicketByID(request.GetId())
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}
	return ToPBGetPickTicketByIdResponse(&ent.PickTicket)
}
