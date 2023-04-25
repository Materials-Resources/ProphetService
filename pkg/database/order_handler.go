package database

import (
	"context"
	"fmt"
	"github.com/uptrace/bun"
)

type OrderHandler struct {
	db *bun.DB
}

func NewOrderHandler(db *bun.DB) *OrderHandler {
	return &OrderHandler{db: db}
}

func (h *OrderHandler) SelectOrder(ctx context.Context, id string) (*Order, error) {
	model := new(Order)
	if err := h.db.NewSelect().Model(model).Where("order_no = ?", id).Scan(ctx); err != nil {
		return nil, fmt.Errorf("could not select order: %v", err)
	}

	return model, nil
}

func (h *OrderHandler) SelectOrderItem(ctx context.Context, orderId string, productId []int32) (*[]OrderLine, error) {
	var model []OrderLine
	query := h.db.NewSelect().Model(&model).Where("order_no = ?", orderId)

	if len(productId) > 0 {
		query.Where("inv_mast_uid in (?)", bun.In(productId))
	}

	if err := query.Scan(ctx); err != nil {
		return nil, fmt.Errorf("could not select order item: %v", err)
	}
	return &model, nil
}
