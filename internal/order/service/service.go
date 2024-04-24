package service

import (
	"context"
	"github.com/materials-resources/s_prophet/internal/order/domain"
)

type Service interface {
	GetOrderById(context.Context, string) (domain.Order, error)
	CreateOrder(context.Context, *domain.Order) error
	UpdateOrder(context.Context, *domain.Order) error
	DeleteOrder(context.Context, string) error

	CreateQuote(context.Context, *domain.Order) error

	GetPickTicketById(context.Context, float64) (domain.PickTicket, error)
}
