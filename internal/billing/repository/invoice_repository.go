package repository

import (
	"context"
	"github.com/materials-resources/s-prophet/internal/billing/domain"
	"github.com/materials-resources/s-prophet/pkg/helpers"
	"github.com/uptrace/bun"
)

func NewInvoiceRepository(db bun.IDB) *InvoiceRepository {
	return &InvoiceRepository{db: db}
}

type InvoiceRepository struct {
	db bun.IDB
}

func (r *InvoiceRepository) GetByOrder(ctx context.Context, orderID string) ([]*domain.Invoice, error) {
	var invoiceHdrRecs []*invoiceHdr
	err := r.db.NewSelect().Model(&invoiceHdrRecs).Where("order_no = ?", orderID).Scan(ctx)
	if err != nil {
		return nil, err
	}

	var invoices []*domain.Invoice
	for _, invoiceHdrRec := range invoiceHdrRecs {
		invoices = append(invoices, &domain.Invoice{
			Id:           invoiceHdrRec.InvoiceNo,
			OrderId:      helpers.GetOptionalValue(invoiceHdrRec.OrderNo, ""),
			AmountPaid:   invoiceHdrRec.AmountPaid,
			Total:        invoiceHdrRec.TotalAmount,
			DateInvoiced: invoiceHdrRec.InvoiceDate,
		})

	}
	return invoices, nil
}
