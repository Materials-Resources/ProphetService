package shipping

import (
	"context"

	"github.com/materials-resources/s_prophet/pkg/internal/core/domain"
	"github.com/materials-resources/s_prophet/pkg/internal/core/port/repository"
	rpc "github.com/materials-resources/s_prophet/proto/shipping/v1alpha0"
	"google.golang.org/grpc"
)

type service struct {
	shippingRepository repository.ShippingRepository
}

func (s service) GetPickTicket(ctx context.Context, request *rpc.GetPickTicketRequest) (
	*rpc.GetPickTicketResponse, error,
) {
	if pickTicket, err := s.shippingRepository.GetPickTicket(request.GetId()); err != nil {
		return nil, err
	} else {
		return domainToPickTicketResponse(pickTicket), err
	}
}

func NewShippingServer(server *grpc.Server, shippingRepository repository.ShippingRepository) {
	shippingServer := &service{shippingRepository: shippingRepository}
	rpc.RegisterShippingServiceServer(
		server,
		shippingServer,
	)
}

func domainToPickTicketResponse(pickTicket *domain.PickTicket) *rpc.GetPickTicketResponse {
	getPickTicketResponse := &rpc.GetPickTicketResponse{
		ShippingDetails: &rpc.GetPickTicketResponse_ShippingDetails{
			Name:                 pickTicket.ShippingAddress.Name,
			LineOne:              pickTicket.ShippingAddress.LineOne,
			LineTwo:              pickTicket.ShippingAddress.LineTwo,
			City:                 pickTicket.ShippingAddress.City,
			State:                pickTicket.ShippingAddress.State,
			PostalCode:           "",
			Country:              "",
			DeliveryInstructions: pickTicket.ShippingAddress.DeliveryInstructions,
		},
		OrderId:            pickTicket.OrderId,
		OrderPurchaseOrder: pickTicket.OrderPurchaseOrder,
		OrderContactName:   pickTicket.OrderContactName,
		PickTicketId:       pickTicket.ID,
	}

	return getPickTicketResponse
}
