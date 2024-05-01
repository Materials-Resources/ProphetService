package api

import (
	"context"
	"fmt"
	"github.com/materials-resources/s-prophet/internal/order/domain"
	"github.com/materials-resources/s-prophet/internal/order/service"

	rpc "github.com/materials-resources/s-prophet/proto/order/v1"
)

func NewOrderApi(service service.Service) OrderApi {
	return OrderApi{service: service}
}

type OrderApi struct {
	service service.Service
}

func (s OrderApi) ListOrdersByCustomer(
	ctx context.Context, request *rpc.ListOrdersByCustomerRequest) (*rpc.ListOrdersByCustomerResponse, error) {
	orders, metadata, err := s.service.ListOrdersByCustomer(
		ctx, request.GetCustomerId(), domain.Filters{
			Direction: domain.PageDirection(request.GetFilters().GetDirection()),
			Cursor:    int(request.GetFilters().GetCursor()),
		})
	if err != nil {
		return nil, err

	}

	res := &rpc.ListOrdersByCustomerResponse{
		Orders: make([]*rpc.BasicOrder, 0, len(orders)),
		Metadata: &rpc.PageMetadata{
			NextCursor: int32(metadata.NextCursor),
			PrevCursor: int32(metadata.PreviousCursor),
		},
	}

	for _, order := range orders {
		res.Orders = append(
			res.Orders, &rpc.BasicOrder{
				Id:            order.Id,
				Taker:         order.Taker,
				PurchaseOrder: order.PurchaseOrder,
			})

	}

	return res, nil
}

func (s OrderApi) ListOrdersByTaker(
	ctx context.Context, request *rpc.ListOrdersByTakerRequest) (*rpc.ListOrdersByTakerResponse, error) {
	// TODO implement me
	panic("implement me")
}

func (s OrderApi) CreateQuote(ctx context.Context, request *rpc.CreateQuoteRequest) (*rpc.CreateQuoteResponse, error) {
	order := &domain.Order{
		Contact: domain.Contact{
			Id: request.GetContactId(),
		},
		ShippingAddress: domain.Address{
			Id: request.GetShippingAddressId(),
		},
		Customer: domain.Customer{
			Id: request.GetCustomerId(),
		},
	}

	order.Items = make([]domain.OrderItem, 0, len(request.GetOrderItems()))
	for _, item := range request.GetOrderItems() {
		order.Items = append(
			order.Items, domain.OrderItem{
				ProductUid:      item.GetProductUid(),
				QuantityOrdered: item.GetQuantityOrdered(),
			})

	}
	err := s.service.CreateQuote(
		ctx, order)

	fmt.Println(order.Id)
	return &rpc.CreateQuoteResponse{
		Id: order.Id,
	}, err
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

	res := &rpc.GetOrderResponse{
		Id:                   order.Id,
		ShippingAddress:      mapAddressToRpcAddress(order.ShippingAddress),
		Contact:              mapContactToRpcContact(order.Contact),
		DeliveryInstructions: order.DeliveryInstructions,
	}

	res.OrderItems = make([]*rpc.GetOrderResponse_OrderItem, 0, len(order.Items))
	for _, item := range order.Items {
		res.OrderItems = append(
			res.OrderItems, &rpc.GetOrderResponse_OrderItem{
				ProductUid:      item.ProductUid,
				QuantityOrdered: item.QuantityOrdered,
			})
	}

	return res, nil
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
