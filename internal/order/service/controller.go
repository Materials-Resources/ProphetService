package service

import (
	"context"
	"github.com/materials-resources/s-prophet/internal/order/data"
	"github.com/materials-resources/s-prophet/internal/order/domain"
)

func NewOrderController(models *data.Models) *OrderController {
	return &OrderController{models: models}
}

type OrderController struct {
	models *data.Models
}

func (c *OrderController) GetOrderCustomer(ctx context.Context, id string) (*domain.Order, error) {
	return c.models.OeHdr.Get(ctx, id)

}

func (c *OrderController) GetShipmentsByOrder(ctx context.Context, id string) ([]*domain.Shipment, error) {
	return c.models.OePickTicket.ListForOrderId(ctx, id)
}

func (c *OrderController) CustomerListOrders(ctx context.Context, id string) ([]*domain.Order, error) {
	if id == "" {
		return nil, nil
	}
	opts := &data.OeHdrListOptions{
		CustomerBranchId: &[]string{id},
	}
	return c.models.OeHdr.List(ctx, opts)
}

func (c *OrderController) ClerkCreateQuote(ctx context.Context, id string) ([]*domain.Order, error) {
	return nil, nil
}
