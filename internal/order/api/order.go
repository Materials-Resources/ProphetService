package api

import (
	"context"
	"fmt"
	"github.com/materials-resources/s-prophet/internal/order/domain"
	"github.com/materials-resources/s-prophet/internal/order/service"
	"google.golang.org/protobuf/types/known/timestamppb"

	rpc "github.com/materials-resources/s-prophet/proto/order/v1"
)

func NewOrderApi(service service.Service) OrderApi {
	return OrderApi{service: service}
}

type OrderApi struct {
	service service.Service
}

func (s OrderApi) ListOrdersByCustomerBranch(
	ctx context.Context, request *rpc.ListOrdersByCustomerBranchRequest) (
	*rpc.ListOrdersByCustomerBranchResponse, error) {
	orders, metadata, err := s.service.ListOrdersByCustomerBranch(
		ctx, request.GetCustomerBranchId(), domain.Filters{
			Direction: mapRpcPageDirectionToDomain(request.GetFilters().GetDirection()),
			Cursor:    int(request.GetFilters().GetCursor()),
		})
	if err != nil {
		return nil, err

	}

	res := &rpc.ListOrdersByCustomerBranchResponse{
		Orders: make([]*rpc.BasicOrder, len(orders)),
		Metadata: &rpc.PageMetadata{
			NextCursor: int32(metadata.NextCursor),
			PrevCursor: int32(metadata.PreviousCursor),
		},
	}

	for i, order := range orders {
		res.Orders[i] = &rpc.BasicOrder{
			Id:            order.Id,
			Status:        "",
			Completed:     order.Completed,
			Taker:         order.Taker,
			PurchaseOrder: order.PurchaseOrder,
			OrderDate:     timestamppb.New(order.RequestedDate),
			RequestedDate: timestamppb.New(order.RequestedDate),
		}
	}

	return res, nil

}

func mapRpcPageDirectionToDomain(direction rpc.PageDirection) domain.PageDirection {
	switch direction {
	case rpc.PageDirection_PAGE_DIRECTION_NEXT:
		return domain.PageDirectionNext
	case rpc.PageDirection_PAGE_DIRECTION_PREVIOUS:
		return domain.PageDirectionPrevious
	default:
		return domain.PageDirectionUnknown
	}
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
	// TODO: get contact id from grpc request context if customer is calling this
	order := &domain.Order{
		Contact: domain.Contact{
			Id: request.GetContactId(),
		},
		CustomerBranchId: request.GetCustomerBranchId(),
		OrderType:        domain.OrderTypeQuote,
		RequestedDate:    request.GetRequestedDate().AsTime(),
		PurchaseOrder:    request.GetPurchaseOrder(),
	}

	order.Items = make([]*domain.OrderItem, len(request.GetOrderItems()))
	for _, item := range request.GetOrderItems() {
		order.Items = append(
			order.Items, &domain.OrderItem{
				ProductUid:    item.GetProductUid(),
				OrderQuantity: item.GetQuantityOrdered(),
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
		CustomerId:           order.CustomerId,
		CustomerName:         order.CustomerName,
		ContactId:            order.ContactId,
		ContactName:          order.ContactName,
		ShippingAddress:      mapAddressToRpcAddress(order.ShippingAddress),
		DeliveryInstructions: order.DeliveryInstructions,
		Taker:                order.Taker,
		PurchaseOrder:        order.PurchaseOrder,
		SubTotal:             0,
		Tax:                  0,
		Total:                0,
		DatePlaced:           timestamppb.New(order.OrderDate),
	}

	res.OrderItems = make([]*rpc.GetOrderResponse_OrderItem, len(order.Items))
	for i, item := range order.Items {
		res.OrderItems[i] = &rpc.GetOrderResponse_OrderItem{
			ProductUid:          item.ProductUid,
			Name:                item.ProductName,
			Sn:                  "",
			Id:                  "",
			CustomerProductSn:   item.CustomerProductSn,
			QuantityOrdered:     item.OrderQuantity,
			QuantityUnit:        item.OrderQuantityUnit,
			CostPerUnit:         item.Price,
			CostUnit:            item.PriceUnit,
			TotalPrice:          item.TotalPrice,
			ShippedQuantity:     item.ShippedQuantity,
			BackOrderedQuantity: item.BackOrderedQuantity,
		}
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
