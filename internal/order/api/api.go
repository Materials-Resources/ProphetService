package api

import (
	"context"
	rpc "github.com/materials-resources/microservices-proto/golang/order"
	"github.com/materials-resources/s-prophet/internal/order/domain"
	"github.com/materials-resources/s-prophet/internal/order/service"
	"github.com/materials-resources/s-prophet/pkg/helpers"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ rpc.OrderServiceServer = &OrderApi{}

type OrderApi struct {
	controller service.Controller
}

func (s *OrderApi) ListOrders(ctx context.Context, request *rpc.ListOrdersRequest) (*rpc.ListOrdersResponse, error) {
	orders, err := s.controller.ListOrders(ctx)
	if err != nil {
		return nil, err
	}
	var pbOrders []*rpc.Order

	for _, order := range orders {
		pbOrders = append(pbOrders, mapOrderWithoutItems(order))
	}

	res := &rpc.ListOrdersResponse{
		Orders: pbOrders,
	}

	return res, nil
}

func mapOrderWithoutItems(order *domain.Order) *rpc.Order {
	return &rpc.Order{
		Id:            order.Id,
		PurchaseOrder: helpers.GetValueOrDefault(order.PurchaseOrder, ""),
		DateCreated:   timestamppb.New(order.DateCreated),
		DateRequested: timestamppb.New(helpers.GetValueOrDefault(order.DateRequested, order.DateCreated)),
		Customer: &rpc.Customer{
			Id:   order.Customer.Id,
			Name: helpers.GetValueOrDefault(order.Customer.Name, ""),
		},
	}
}

func (s *OrderApi) GetOrder(ctx context.Context, request *rpc.GetOrderRequest) (*rpc.GetOrderResponse, error) {
	order, err := s.controller.GetOrder(ctx, request.GetId())

	if err != nil {
		return nil, err

	}

	res := &rpc.GetOrderResponse{
		Order:      mapOrder(order),
		OrderItems: make([]*rpc.OrderItem, len(order.Items)),
	}

	for i, item := range order.Items {
		res.OrderItems[i] = mapOrderItem(item)
	}
	return res, nil
}

func mapOrder(order *domain.Order) *rpc.Order {
	return &rpc.Order{
		Id:            order.Id,
		PurchaseOrder: helpers.GetValueOrDefault(order.PurchaseOrder, ""),
		DateCreated:   timestamppb.New(order.DateCreated),
		DateRequested: timestamppb.New(helpers.GetValueOrDefault(order.DateRequested, order.DateCreated)),
		Customer: &rpc.Customer{
			Id:   order.Customer.Id,
			Name: helpers.GetValueOrDefault(order.Customer.Name, ""),
		},
	}
}

func mapOrderItem(item *domain.OrderItem) *rpc.OrderItem {
	return &rpc.OrderItem{
		Id:                  item.Id,
		ProductUid:          item.ProductUid,
		ProductSn:           item.ProductSn,
		ProductName:         item.ProductName,
		CustomerProductSn:   item.CustomerProductSn,
		QuantityOrdered:     item.OrderQuantity,
		QuantityUnit:        item.OrderQuantityUnit,
		TotalPrice:          item.TotalPrice,
		ShippedQuantity:     item.ShippedQuantity,
		BackOrderedQuantity: item.CalculateBackOrderedQuantity(),
	}
}

func (s *OrderApi) CreateOrder(ctx context.Context, request *rpc.CreateOrderRequest) (*rpc.CreateOrderResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *OrderApi) GetQuote(ctx context.Context, request *rpc.GetQuoteRequest) (*rpc.GetQuoteResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *OrderApi) CreateQuote(ctx context.Context, request *rpc.CreateQuoteRequest) (*rpc.CreateQuoteResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *OrderApi) ListShipmentsByOrder(ctx context.Context, request *rpc.ListShipmentsByOrderRequest) (*rpc.ListShipmentsByOrderResponse, error) {
	shipments, err := s.controller.ListShipmentsByOrder(ctx, request.GetOrderId())
	if err != nil {
		return nil, err
	}

	res := &rpc.ListShipmentsByOrderResponse{
		Shipments: make([]*rpc.Shipment, 0, len(shipments)),
	}
	for _, shipment := range shipments {
		res.Shipments = append(res.Shipments, mapShipment(shipment))
	}

	return res, nil
}

func mapShipment(shipment *domain.Shipment) *rpc.Shipment {
	return &rpc.Shipment{
		Id:                 shipment.Id,
		CarrierName:        shipment.CarrierName,
		CarrierTracking:    helpers.GetValueOrDefault(shipment.CarrierTracking, ""),
		OrderId:            shipment.OrderId,
		OrderPurchaseOrder: helpers.GetValueOrDefault(shipment.OrderPurchaseOrder, ""),
		Customer: &rpc.Customer{
			Id:   shipment.Customer.Id,
			Name: helpers.GetValueOrDefault(shipment.Customer.Name, ""),
		},
		ShippingAddress: &rpc.Address{
			Id:         shipment.ShippingAddress.Id,
			Name:       helpers.GetValueOrDefault(shipment.ShippingAddress.Name, ""),
			LineOne:    helpers.GetValueOrDefault(shipment.ShippingAddress.LineOne, ""),
			LineTwo:    helpers.GetValueOrDefault(shipment.ShippingAddress.LineTwo, ""),
			City:       helpers.GetValueOrDefault(shipment.ShippingAddress.City, ""),
			State:      helpers.GetValueOrDefault(shipment.ShippingAddress.State, ""),
			PostalCode: helpers.GetValueOrDefault(shipment.ShippingAddress.PostalCode, ""),
			Country:    helpers.GetValueOrDefault(shipment.ShippingAddress.Country, ""),
		},
	}
}

func (s *OrderApi) GetShipment(ctx context.Context, request *rpc.GetShipmentRequest) (*rpc.GetShipmentResponse, error) {
	shipment, err := s.controller.GetShipment(ctx, request.GetId())
	if err != nil {
		return nil, err
	}

	res := &rpc.GetShipmentResponse{
		Shipment: mapShipment(shipment),
	}

	return res, nil
}

func NewOrderApi(controller service.Controller) *OrderApi {
	return &OrderApi{controller: controller}
}
