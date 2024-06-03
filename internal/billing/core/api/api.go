package api

import (
	"context"
	"github.com/materials-resources/s-prophet/infrastructure/data"
	"github.com/materials-resources/s-prophet/internal/billing/core/domain"
)

type BillingAdapter struct {
	models *data.Models
}

func NewBillingAdapter(models data.Models) *BillingAdapter {
	return &BillingAdapter{
		models: &models,
	}
}

func (a *BillingAdapter) GetInvoicesByOrderID(orderID string) ([]*domain.Invoice, error) {
	invoiceHdrRecords, err := a.models.InvoiceHdr.GetAll(context.Background(), data.InvoiceHdrGetAllParams{OrderNo: &[]string{orderID}})
	if err != nil {
		return nil, err
	}
	invoices := make([]*domain.Invoice, len(invoiceHdrRecords))
	for i, record := range invoiceHdrRecords {
		invoices[i] = &domain.Invoice{
			Id:        record.InvoiceNo,
			OrderId:   record.OrderNo.String,
			Total:     record.TotalAmount,
			CreatedAt: record.DateCreated,
		}
	}
	return invoices, nil
}
