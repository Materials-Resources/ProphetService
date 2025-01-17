package service

import (
	"context"
	orderv1 "github.com/materials-resources/s-prophet/internal/grpc/order"
	"github.com/materials-resources/s-prophet/internal/order/domain"
	"github.com/materials-resources/s-prophet/internal/order/repository"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type OrderService interface {
	// GetOrder returns an Order with its details
	GetOrder(ctx context.Context, req *orderv1.GetOrderRequest) (*orderv1.GetOrderResponse, error)
	CreateOrder(ctx context.Context)
	ListOrders(ctx context.Context, req *orderv1.ListOrdersRequest) (*orderv1.ListOrdersResponse, error)
	ListQuotes(ctx context.Context, req *orderv1.ListQuotesRequest) (*orderv1.ListQuotesResponse, error)
	CreateQuote(ctx context.Context, req *orderv1.CreateQuoteRequest) (*orderv1.CreateQuoteResponse, error)
	ListShipmentsByOrder(ctx context.Context, orderId string) ([]*domain.Shipment, error)
	GetShipment(ctx context.Context, id string) (*domain.Shipment, error)
}

var _ OrderService = &Order{}

type Order struct {
	repository *repository.Repository
}

func (s *Order) ListQuotes(ctx context.Context, req *orderv1.ListQuotesRequest) (*orderv1.ListQuotesResponse, error) {
	params := repository.ListQuotesParams{
		BranchId: req.GetBranchId(),
		Page:     int(req.GetPage()),
		PageSize: int(req.GetPageSize()),
	}
	quotes, err := s.repository.Order.ListQuotes(ctx, params)
	if err != nil {
		return nil, err
	}

	response := &orderv1.ListQuotesResponse{}

	for _, quote := range quotes {
		response.Quotes = append(response.Quotes, &orderv1.QuoteSummary{
			Id: quote.Id,
			Contact: &orderv1.Contact{
				Id: quote.ContactId,
			},
			Branch: &orderv1.Branch{
				Id: quote.BranchId,
			},
			PurchaseOrder: quote.PurchaseOrder,
			Status:        convertQuoteStatus(quote.Status),
			DateCreated:   timestamppb.New(quote.DateCreated),
			DateExpires:   timestamppb.New(quote.DateExpires),
		})
	}

	return response, nil
}

func (s *Order) CreateQuote(ctx context.Context, req *orderv1.CreateQuoteRequest) (*orderv1.CreateQuoteResponse, error) {

	createQuoteReq := repository.CreateQuoteParams{
		PurchaseOrder:        req.GetPurchaseOrder(),
		ContactId:            req.GetContactId(),
		BranchId:             req.GetBranchId(),
		RequestDate:          req.GetRequestedDate().AsTime(),
		DeliveryInstructions: req.GetDeliveryInstructions(),
	}

	for _, item := range req.GetItems() {
		createQuoteReq.Items = append(createQuoteReq.Items, repository.CreateQuoteItemParams{
			ProductId: item.GetProductId(),
			Quantity:  item.GetQuantity(),
		})
	}
	quoteId, err := s.repository.Order.CreateQuote(ctx, &createQuoteReq)
	if err != nil {
		return nil, err
	}
	return &orderv1.CreateQuoteResponse{
		Id: *quoteId,
	}, nil
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

func (s *Order) GetOrder(ctx context.Context, req *orderv1.GetOrderRequest) (*orderv1.GetOrderResponse, error) {

	order, err := s.repository.Order.GetOrder(ctx, req.GetId())

	if err != nil {
		return nil, err
	}

	response := &orderv1.GetOrderResponse{
		Order: &orderv1.Order{
			Id:            order.Id,
			Status:        convertOrderStatus(order.Status),
			PurchaseOrder: order.PurchaseOrder,
			DateCreated:   timestamppb.New(order.DateCreated),
			DateRequested: timestamppb.New(order.DateRequested),
			BranchId:      order.BranchId,
			ContactId:     order.ContactId,
			ShippingAddress: &orderv1.Address{
				Id:         order.ShippingAddress.Id,
				Name:       order.ShippingAddress.Name,
				LineOne:    order.ShippingAddress.LineOne,
				LineTwo:    order.ShippingAddress.LineTwo,
				City:       order.ShippingAddress.City,
				State:      order.ShippingAddress.State,
				PostalCode: order.ShippingAddress.PostalCode,
				Country:    order.ShippingAddress.Country,
			},
		},
	}

	for _, item := range order.Items {
		response.Order.OrderItems = append(response.Order.OrderItems, &orderv1.Order_Item{
			Id:                  "",
			ProductId:           item.ProductId,
			ProductSn:           item.ProductSn,
			ProductName:         item.ProductName,
			CustomerProductSn:   item.CustomerProductSn,
			OrderedQuantity:     item.OrderedQuantity,
			UnitPrice:           item.UnitPrice,
			UnitType:            item.UnitType,
			TotalPrice:          item.TotalPrice,
			ShippedQuantity:     item.ShippedQuantity,
			BackOrderedQuantity: 0,
		})
	}

	return response, nil

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

		response.Orders = append(response.Orders, &orderv1.OrderSummary{
			Id:            order.Id,
			Status:        convertOrderStatus(order.Status),
			PurchaseOrder: order.PurchaseOrder,
			DateCreated:   timestamppb.New(order.DateCreated),
			DateRequested: timestamppb.New(order.DateRequested),
			BranchId:      order.BranchId,
			ContactId:     order.ContactId,
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

func convertOrderStatus(status domain.OrderStatus) orderv1.OrderStatus {
	switch status {
	case domain.OrderStatusCompleted:
		return orderv1.OrderStatus_ORDER_STATUS_COMPLETED
	case domain.OrderStatusPendingApproval:
		return orderv1.OrderStatus_ORDER_STATUS_PENDING_APPROVAL
	case domain.OrderStatusApproved:
		return orderv1.OrderStatus_ORDER_STATUS_APPROVED
	case domain.OrderStatusCancelled:
		return orderv1.OrderStatus_ORDER_STATUS_CANCELLED
	default:
		return orderv1.OrderStatus_ORDER_STATUS_UNSPECIFIED
	}
}

func convertQuoteStatus(status domain.QuoteStatus) orderv1.QuoteStatus {
	switch status {
	case domain.QuoteStatusApproved:
		return orderv1.QuoteStatus_QUOTE_STATUS_APPROVED
	case domain.QuoteStatusCancelled:
		return orderv1.QuoteStatus_QUOTE_STATUS_CANCELLED
	case domain.QuoteStatusPendingApproval:
		return orderv1.QuoteStatus_QUOTE_STATUS_PENDING_APPROVAL
	case domain.QuoteStatusExpired:
		return orderv1.QuoteStatus_QUOTE_STATUS_EXPIRED
	default:
		return orderv1.QuoteStatus_QUOTE_STATUS_UNSPECIFIED

	}
}
