package data

import (
	"context"
	"github.com/materials-resources/s-prophet/internal/billing/domain"
	"github.com/uptrace/bun"
)

type InvoiceModel struct {
	db bun.IDB
}

func (m *InvoiceModel) GetByOrder(ctx context.Context, orderId string) ([]*domain.Invoice, error) {
	var recs []*invoiceHdr

	err := m.db.NewSelect().Model(&recs).Where("order_no = ?", orderId).Scan(ctx)

	if err != nil {
		return nil, err
	}

	return nil, nil
}
