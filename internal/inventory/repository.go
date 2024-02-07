package inventory

import (
	"github.com/materials-resources/s_prophet/internal/inventory/domain"
)

type Repository interface {
	FindReceiptById(id string) (*domain.ValidatedInventoryReceipt, error)
}
