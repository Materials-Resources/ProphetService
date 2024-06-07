package service

import (
	"context"
	"github.com/materials-resources/s-prophet/internal/order/domain"
)

type Service interface {
	GetOrderById(context.Context, string) (domain.Order, error)
	ListOrdersByCustomer(context.Context, float64, domain.Filters) ([]*domain.Order, domain.Metadata, error)
	ListOrdersByCustomerBranch(
		ctx context.Context, customerBranchId float64,
		filters domain.Filters) (
		[]*domain.Order, domain.Metadata,
		error)

	CreateQuote(context.Context, *domain.Order) error

	GetPickTicketById(context.Context, float64) (domain.PickTicket, error)
}
