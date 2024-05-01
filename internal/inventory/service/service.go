package service

import (
	"context"
	"github.com/materials-resources/s-prophet/internal/inventory/domain"
)

type Service interface {
	GetProductStock(context.Context, []int32) ([]domain.ProductStock, error)
	GetInventoryReceiptById(context.Context, float64) (domain.InventoryReceipt, error)
}
