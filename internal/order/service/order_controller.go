package service

import (
	"context"
	"github.com/materials-resources/s-prophet/internal/order/domain"
	"github.com/materials-resources/s-prophet/internal/order/repository"
)

type Controller interface {
	// GetOrder returns an Order with its details
	GetOrder(ctx context.Context, id string) (*domain.Order, error)
	CreateOrder(ctx context.Context)
	ListOrders(ctx context.Context) ([]*domain.Order, error)
	ListShipmentsByOrder(ctx context.Context, orderId string) ([]*domain.Shipment, error)
	GetShipment(ctx context.Context, id string) (*domain.Shipment, error)
}

var _ Controller = &OrderService{}

type OrderService struct {
	repository *repository.Repository
}

func (s *OrderService) CreateOrder(ctx context.Context) {
	//TODO implement me
	panic("implement me")
}

func (s *OrderService) GetShipment(ctx context.Context, id string) (*domain.Shipment, error) {
	shipment, err := s.repository.Shipment.GetShipment(ctx, id)
	if err != nil {
		return nil, err
	}
	return shipment, nil
}

func (s *OrderService) GetOrder(ctx context.Context, id string) (*domain.Order, error) {

	order, err := s.repository.Order.GetOrder(ctx, id)

	if err != nil {
		return nil, err
	}

	return order, nil

}

func (s *OrderService) ListOrders(ctx context.Context) ([]*domain.Order, error) {
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

func (s *OrderService) ListShipmentsByOrder(ctx context.Context, orderId string) ([]*domain.Shipment, error) {
	shipments, err := s.repository.Shipment.ListShipmentsByOrder(ctx, orderId)
	if err != nil {
		return nil, err
	}

	return shipments, nil
}
func NewOrderController(repo *repository.Repository) *OrderService {
	return &OrderService{repository: repo}
}
