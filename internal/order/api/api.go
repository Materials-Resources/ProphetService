package api

import (
	"connectrpc.com/connect"
	"context"
	"github.com/materials-resources/s-prophet/internal/grpc/order"
	"github.com/materials-resources/s-prophet/internal/grpc/order/orderconnect"
	"github.com/materials-resources/s-prophet/internal/order/service"
)

func NewOrderServiceHandler(svc service.OrderService) *OrderServiceHandler {
	return &OrderServiceHandler{
		service: svc,
	}
}

type OrderServiceHandler struct {
	service service.OrderService
}

func (h OrderServiceHandler) ListOrders(ctx context.Context, req *connect.Request[order.ListOrdersRequest]) (*connect.Response[order.ListOrdersResponse], error) {
	res, err := h.service.ListOrders(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(res), nil
}

func (h OrderServiceHandler) GetOrder(ctx context.Context, req *connect.Request[order.GetOrderRequest]) (*connect.Response[order.GetOrderResponse], error) {
	res, err := h.service.GetOrder(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(res), nil
}

func (h OrderServiceHandler) CreateOrder(ctx context.Context, req *connect.Request[order.CreateOrderRequest]) (*connect.Response[order.CreateOrderResponse], error) {
	//TODO implement me
	panic("implement me")
}

func (h OrderServiceHandler) GetQuote(ctx context.Context, req *connect.Request[order.GetQuoteRequest]) (*connect.Response[order.GetQuoteResponse], error) {
	//TODO implement me
	panic("implement me")
}

func (h OrderServiceHandler) CreateQuote(ctx context.Context, req *connect.Request[order.CreateQuoteRequest]) (*connect.Response[order.CreateQuoteResponse], error) {
	res, err := h.service.CreateQuote(ctx, req.Msg)

	if err != nil {
		return nil, err
	}

	return connect.NewResponse(res), nil
}

func (h OrderServiceHandler) ListShipmentsByOrder(ctx context.Context, req *connect.Request[order.ListShipmentsByOrderRequest]) (*connect.Response[order.ListShipmentsByOrderResponse], error) {
	//TODO implement me
	panic("implement me")
}

func (h OrderServiceHandler) GetShipment(ctx context.Context, req *connect.Request[order.GetShipmentRequest]) (*connect.Response[order.GetShipmentResponse], error) {
	//TODO implement me
	panic("implement me")
}

var _ orderconnect.OrderServiceHandler = (*OrderServiceHandler)(nil)
