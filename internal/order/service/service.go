package service

import (
	"context"
	"github.com/materials-resources/s_prophet/internal/order/domain"
)

type Service interface {
	GetOrderById(context.Context, string) (domain.Order, error)

	GetPickTicketById(context.Context, float64) (domain.PickTicket, error)
}
