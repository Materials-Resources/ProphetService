package service

import (
	"context"
	"github.com/materials-resources/s-prophet/internal/order/domain"
)

func (c *OrderController) ClerkGetOrder(ctx context.Context, id string) (*domain.Order, error) {
	return c.models.OeHdr.Get(ctx, id)
}

func (c *OrderController) ClerkGetShipment(ctx context.Context, id string) (*domain.Shipment, error) {
	return c.models.OePickTicket.Get(ctx, id)
}
