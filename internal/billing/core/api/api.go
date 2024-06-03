package api

import "github.com/materials-resources/s-prophet/internal/billing/core/domain"

type BillingAdapter struct {
}

func NewBillingAdapter() *BillingAdapter {
	return &BillingAdapter{}
}

func (a *BillingAdapter) GetInvoicesByOrderID(orderID string) ([]*domain.Invoice, error) {
	//TODO implement me
	panic("implement me")
}
