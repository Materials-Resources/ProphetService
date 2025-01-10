package service

import (
	"context"
	orderv1 "github.com/materials-resources/s-prophet/internal/grpc/order"
	"github.com/materials-resources/s-prophet/internal/order/domain"
	"github.com/materials-resources/s-prophet/internal/order/repository"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type OrderService interface {
	// GetOrder returns an Order with its details
	GetOrder(ctx context.Context, id string) (*domain.Order, error)
	CreateOrder(ctx context.Context)
	ListOrders(ctx context.Context, req *orderv1.ListOrdersRequest) (*orderv1.ListOrdersResponse, error)
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

func (s *Order) ListOrders(ctx context.Context, req *orderv1.ListOrdersRequest) (*orderv1.ListOrdersResponse, error) {
	params := repository.ListOrdersParams{
		BranchId: req.GetBranchId(),
		Page:     int(req.GetPage()),
		PageSize: int(req.GetPageSize()),
	}

	orders, err := s.repository.Order.ListOrders(ctx, params)
	if err != nil {
		return nil, err
	}

	response := &orderv1.ListOrdersResponse{}

	for _, order := range orders {
		response.Orders = append(response.Orders, &orderv1.Order{
			Id:            order.Id,
			Status:        0,
			PurchaseOrder: order.PurchaseOrder,
			DateCreated:   timestamppb.New(order.DateCreated),
		})
	}

	return response, nil
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
