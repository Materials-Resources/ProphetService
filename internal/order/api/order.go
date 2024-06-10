package api

import (
	"context"
	"fmt"
	"github.com/materials-resources/s-prophet/internal/order/domain"
	"github.com/materials-resources/s-prophet/internal/order/service"
	"strconv"

	rpc "github.com/materials-resources/microservices-proto/golang/order"
)

func NewOrderApi(controller service.OrderController) OrderApi {
	return OrderApi{controller: controller}
}

type OrderApi struct {
	controller service.OrderController
}

func (s OrderApi) ClerkGetShipmentsForOrder(ctx context.Context, request *rpc.ClerkGetShipmentsForOrderRequest) (*rpc.ClerkGetShipmentsForOrderResponse, error) {
	panic("implement me")
}

func (s OrderApi) ClerkGetOrder(ctx context.Context, request *rpc.ClerkGetOrderRequest) (*rpc.ClerkGetOrderResponse, error) {
	order, err := s.controller.ClerkGetOrder(ctx, request.GetId())

	if err != nil {
		return nil, err

	}
	return CreateClerkGetOrderResponse(order), nil

}

func (s OrderApi) ClerkCreateOrder(ctx context.Context, request *rpc.ClerkCreateOrderRequest) (*rpc.ClerkCreateOrderResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s OrderApi) ClerkCreateQuote(ctx context.Context, request *rpc.ClerkCreateQuoteRequest) (*rpc.ClerkCreateQuoteResponse, error) {
	// TODO: get contact id from grpc request context if customer is calling this
	order := &domain.Order{
		//CustomerBranchId: request.GetCustomerBranchId(),
		PurchaseOrder: request.GetPurchaseOrder(),
	}

	order.Items = make([]*domain.OrderItem, len(request.GetOrderItems()))
	for _, item := range request.GetOrderItems() {
		order.Items = append(
			order.Items, &domain.OrderItem{
				ProductUid:    strconv.Itoa(int(item.GetProductUid())),
				OrderQuantity: item.GetQuantityOrdered(),
			})

	}

	//err := s.controller.ClerkCreateQuote(ctx, order)

	fmt.Println(order.Id)
	return &rpc.ClerkCreateQuoteResponse{
		Id: order.Id,
	}, nil
}

func (s OrderApi) ClerkGetShipment(ctx context.Context, request *rpc.ClerkGetShipmentRequest) (*rpc.ClerkGetShipmentResponse, error) {
	shipment, err := s.controller.ClerkGetShipment(ctx, request.GetId())
	if err != nil {
		return nil, err
	}

	return CreateClerkGetShipmentResponse(shipment), nil
}

func (s OrderApi) CustomerGetShipmentsForOrder(ctx context.Context, request *rpc.CustomerGetShipmentsForOrderRequest) (*rpc.CustomerGetShipmentsForOrderResponse, error) {
	shipments, err := s.controller.GetShipmentsByOrder(ctx, request.GetOrderId())
	if err != nil {
		return nil, err
	}

	return CreateCustomerGetShipmentsForOrderResponse(shipments), nil
}

func (s OrderApi) CustomerGetOrder(ctx context.Context, request *rpc.CustomerGetOrderRequest) (*rpc.CustomerGetOrderResponse, error) {
	order, err := s.controller.GetOrderCustomer(ctx, request.GetId())
	if err != nil {
		return nil, err

	}

	return CreateCustomerGetOrderResponse(order), nil
}

func (s OrderApi) CustomerGetQuote(ctx context.Context, request *rpc.CustomerGetQuoteRequest) (*rpc.CustomerGetQuoteResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s OrderApi) CustomerListOrders(ctx context.Context, request *rpc.CustomerListOrdersRequest) (*rpc.CustomerListOrdersResponse, error) {
	orders, err := s.controller.CustomerListOrders(ctx, request.GetCustomerBranchId())
	if err != nil {
		return nil, err

	}

	return CreateCustomerListOrdersResponse(orders), nil
}
