package service

import (
	"context"
	orderv1 "github.com/materials-resources/s-prophet/internal/grpc/order"
	"github.com/materials-resources/s-prophet/internal/order/domain"
	"github.com/materials-resources/s-prophet/internal/order/repository"
	"time"
)

type OrderService interface {
	// GetOrder returns an Order with its details
	GetOrder(ctx context.Context, id string) (*domain.Order, error)
	CreateOrder(ctx context.Context)
	ListOrders(ctx context.Context) ([]*domain.Order, error)
	CreateQuote(ctx context.Context, req *orderv1.CreateQuoteRequest) (*orderv1.CreateQuoteResponse, error)
	ListShipmentsByOrder(ctx context.Context, orderId string) ([]*domain.Shipment, error)
	GetShipment(ctx context.Context, id string) (*domain.Shipment, error)
}

var _ OrderService = &Order{}

type Order struct {
	repository *repository.Repository
}

func (s *Order) CreateQuote(ctx context.Context, req *orderv1.CreateQuoteRequest) (*orderv1.CreateQuoteResponse, error) {

	createQuoteReq := repository.CreateQuoteParams{
		PurchaseOrder:        req.GetPurchaseOrder(),
		CustomerId:           req.GetCustomerId(),
		ContactId:            req.GetContactId(),
		RequestDate:          time.Time{},
		DeliveryInstructions: req.GetDeliveryInstructions(),
	}

	for _, item := range req.GetItems() {
		createQuoteReq.Items = append(createQuoteReq.Items, repository.CreateQuoteItemParams{
			ProductId: item.GetProductId(),
			//TODO update proto to int
			Quantity: int(item.GetQuantity()),
		})
	}
	s.repository.Order.CreateQuote(ctx, &createQuoteReq)
	return &orderv1.CreateQuoteResponse{}, nil
}

func (s *Order) CreateOrder(ctx context.Context) {
	//TODO implement me
	panic("implement me")
}

func (s *Order) GetShipment(ctx context.Context, id string) (*domain.Shipment, error) {
	shipment, err := s.repository.Shipment.GetShipment(ctx, id)
	if err != nil {
		return nil, err
	}
	return shipment, nil
}

func (s *Order) GetOrder(ctx context.Context, id string) (*domain.Order, error) {

	order, err := s.repository.Order.GetOrder(ctx, id)

	if err != nil {
		return nil, err
	}

	return order, nil

}

func (s *Order) ListOrders(ctx context.Context) ([]*domain.Order, error) {
	var params repository.ListOrdersParams

	role := "customer"
	// TODO  logic if a customer is making this request
	if role == "customer" {
		params.OnlyActiveOrders = true
	}

	orders, err := s.repository.Order.ListOrders(ctx, params)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (s *Order) ListShipmentsByOrder(ctx context.Context, orderId string) ([]*domain.Shipment, error) {
	shipments, err := s.repository.Shipment.ListShipmentsByOrder(ctx, orderId)
	if err != nil {
		return nil, err
	}

	return shipments, nil
}
func NewOrderService(repo *repository.Repository) *Order {
	return &Order{repository: repo}
}
