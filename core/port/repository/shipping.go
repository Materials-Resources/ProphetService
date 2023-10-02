package repository

import (
	"context"

	"github.com/materials-resources/s_prophet/core/domain"
)

type ShippingRepository interface {
	ReadPickTicket(ctx context.Context, id string) (*domain.PickTicket, error)
}
