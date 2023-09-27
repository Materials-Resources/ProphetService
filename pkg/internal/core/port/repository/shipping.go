package repository

import (
	"github.com/materials-resources/s_prophet/pkg/internal/core/domain"
)

type ShippingRepository interface {
	GetPickTicket(id string) (*domain.PickTicket, error)
}
