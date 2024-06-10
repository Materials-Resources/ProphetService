package api

import (
	"context"
	"github.com/materials-resources/s-prophet/internal/billing/core/domain"
	"github.com/materials-resources/s-prophet/pkg/models"
)

type BillingAdapter struct {
	m *models.Models
}

func NewBillingAdapter(models models.Models) *BillingAdapter {
	return &BillingAdapter{
		m: &models,
	}
}

func (a *BillingAdapter) GetInvoicesByOrderID(orderID string) ([]*domain.Invoice, error) {
	invoiceHdrRecords, err := a.m.InvoiceHdr.GetAll(context.Background(), models.InvoiceHdrGetAllParams{OrderNo: &[]string{orderID}})
	if err != nil {
		return nil, err
	}
	invoices := make([]*domain.Invoice, len(invoiceHdrRecords))
	for i, record := range invoiceHdrRecords {
		invoices[i] = &domain.Invoice{
			Id:         record.InvoiceNo,
			OrderId:    record.OrderNo.String,
			Total:      record.TotalAmount,
			CreatedAt:  record.DateCreated,
			AmountPaid: record.AmountPaid,
		}
	}
	return invoices, nil
}
