package repository

import (
	"context"

	"github.com/materials-resources/s_prophet/pkg/internal/core/domain"
)

type ShippingRepository interface {
	GetPickTicket(ctx context.Context, id string) (*domain.PickTicket, error)
}
