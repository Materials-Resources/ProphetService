package repositories

import (
	"github.com/materials-resources/s_prophet/pkg/domain/entities"
)

type InventoryRepository interface {
	FindReceiptById(id string) (*entities.ValidatedInventoryReceipt, error)
}
