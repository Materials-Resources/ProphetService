package service

import (
	"context"
	"errors"
	"github.com/materials-resources/s-prophet/internal/billing/data"
	"github.com/materials-resources/s-prophet/internal/billing/domain"
)

var (
	ErrNotAuthorized    = errors.New("not authorized")
	ErrResourceNotFound = errors.New("resource not found")
)

type Billing struct {
	model *data.Models
}

func NewBilling(model *data.Models) *Billing {
	return &Billing{model: model}
}

func (s *Billing) GetInvoicesByOrder(ctx context.Context, orderId string) ([]*domain.Invoice, error) {

	invoices, err := s.model.Invoices.GetByOrder(ctx, orderId)

	if err != nil {
		return nil, err
	}

	return invoices, nil
}
