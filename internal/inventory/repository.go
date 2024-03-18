package inventory

import (
	"context"

	"github.com/materials-resources/s_prophet/internal/inventory/domain"
)

type Repository interface {
	SelectProductStockById(ctx context.Context, id []int32) ([]*domain.ProductStock, error)
	FindReceiptById(id string) (*domain.ValidatedInventoryReceipt, error)
}
