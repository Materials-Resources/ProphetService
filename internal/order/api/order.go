package api

import (
	"context"
	"github.com/materials-resources/s_prophet/internal/order/domain"
	"github.com/materials-resources/s_prophet/internal/order/service"

	rpc "github.com/materials-resources/s_prophet/proto/order/v1"
)

func NewOrderApi(service service.Service) OrderApi {
	return OrderApi{service: service}
}

type OrderApi struct {
	service service.Service
}

func (s OrderApi) CreateQuote(ctx context.Context, request *rpc.CreateQuoteRequest) (*rpc.CreateQuoteResponse, error) {
	s.service.CreateQuote(
		ctx, &domain.Order{
			ShippingAddress: domain.Address{
				Id: 100023,
			},
			Customer: domain.Customer{
				Id: 100023,
			},
		})
	return &rpc.CreateQuoteResponse{}, nil
}

func (s OrderApi) CreateOrder(ctx context.Context, request *rpc.CreateOrderRequest) (
	*rpc.CreateOrderResponse,
	error,
) {
	return nil, nil
}

func (s OrderApi) GetOrder(ctx context.Context, request *rpc.GetOrderRequest) (*rpc.GetOrderResponse, error) {
	order, err := s.service.GetOrderById(ctx, request.GetId())
	if err != nil {
		return nil, err
	}

	return &rpc.GetOrderResponse{
		Id:                   order.Id,
		ShippingAddress:      mapAddressToRpcAddress(order.ShippingAddress),
		Contact:              mapContactToRpcContact(order.Contact),
		DeliveryInstructions: order.DeliveryInstructions,
	}, nil
}

func (s OrderApi) GetPickTicketById(
	ctx context.Context, request *rpc.GetPickTicketByIdRequest,
) (*rpc.GetPickTicketByIdResponse, error) {

	pickTicket, err := s.service.GetPickTicketById(ctx, request.GetId())
	if err != nil {
		return nil, err
	}

	res := &rpc.GetPickTicketByIdResponse{
		Id:                   pickTicket.Id,
		OrderId:              pickTicket.OrderId,
		DeliveryInstructions: pickTicket.DeliveryInstructions,
		OrderPurchaseOrder:   pickTicket.OrderPurchaseOrder,
		OrderContact:         mapContactToRpcContact(pickTicket.OrderContact),
		ShippingAddress:      mapAddressToRpcAddress(pickTicket.ShippingAddress),
	}

	return res, nil
}

func mapAddressToRpcAddress(address domain.Address) *rpc.Address {
	return &rpc.Address{
		Id:         address.Id,
		Name:       address.Name,
		LineOne:    address.LineOne,
		LineTwo:    address.LineTwo,
		City:       address.City,
		State:      address.State,
		PostalCode: address.PostalCode,
	}
}

func mapContactToRpcContact(contact domain.Contact) *rpc.CustomerContact {
	return &rpc.CustomerContact{
		Id:    contact.Id,
		Name:  contact.Name,
		Title: contact.Title,
	}
}
