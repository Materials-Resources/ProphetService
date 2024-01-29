package grpc

import (
	"context"
	"fmt"

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

func (s orderService) GetOrder(ctx context.Context, request *rpc.GetOrderRequest) (*rpc.GetOrderResponse, error) {
	o, err := s.repo.FindOrderById(request.GetId())
	fmt.Println(err)
	return ToPBGetOrderResponse(&o.Order)
}

func (s orderService) GetPickTicketById(ctx context.Context, request *rpc.GetPickTicketByIdRequest,
) (*rpc.GetPickTicketByIdResponse, error) {
	ent, err := s.repo.FindPickTicketByID(request.GetId())
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}
	return ToPBGetPickTicketByIdResponse(&ent.PickTicket)
}
