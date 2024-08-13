package data

import (
	"context"
	"github.com/uptrace/bun"
)

type InvoiceHdrModel struct {
	bun bun.IDB
}

func NewInvoiceHdrModel(bun bun.IDB) InvoiceHdrModel {
	return InvoiceHdrModel{
		bun: bun,
	}
}

func (m *InvoiceHdrModel) Get(ctx context.Context, invoiceNo string) (*InvoiceHdr, error) {
	var rec InvoiceHdr
	err := m.bun.NewSelect().Model(&rec).Where("invoice_no = ?", invoiceNo).Scan(ctx)
	if err != nil {
		return nil, err

	}
	return &rec, err
}
