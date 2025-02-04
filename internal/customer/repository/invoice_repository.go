package repository

import (
	"context"
	"github.com/materials-resources/s-prophet/internal/customer/domain"
	"github.com/materials-resources/s-prophet/pkg/helpers"
	"github.com/uptrace/bun"
	"strconv"
)

type InvoiceRepository struct {
	db bun.IDB
}

func NewInvoiceRepository(db bun.IDB) *InvoiceRepository {
	return &InvoiceRepository{
		db: db,
	}
}

func (r *InvoiceRepository) GetRecentPurchasesByBranch(ctx context.Context, branchID string,
	limit int32) ([]*domain.Purchase, error) {
	var invoiceLines []*InvoiceLine

	const (
		invoiceAdjustmentExclude = "c"
		invoiceClassExclude      = "'Z','FINANCE'"
		invoiceLineTypesExclude  = "4,928,929,3226"
		taxItemRequired          = "N"
	)

	err := r.db.NewSelect().Model(&invoiceLines).
		Relation("InvoiceHdr", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Where("ship_to_id = ?", branchID).
				Where("invoice_adjustment_type <> ?", invoiceAdjustmentExclude).
				Where("invoice_class NOT IN (" + invoiceClassExclude + ")")
		}).
		Where("invoice_line_type NOT IN ("+invoiceLineTypesExclude+")").
		Where("tax_item = ?", taxItemRequired).
		Limit(int(limit)).
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	var purchases []*domain.Purchase

	for _, invoiceLine := range invoiceLines {
		purchase := &domain.Purchase{
			ProductId:   strconv.Itoa(int(helpers.GetOptionalValue(invoiceLine.InvMastUid, 0))),
			ProductSn:   invoiceLine.ItemId,
			ProductName: &invoiceLine.ItemDesc,
			UnitType:    invoiceLine.UnitOfMeasure,
		}
		purchases = append(purchases, purchase)
	}
	return purchases, nil
}

func (r *InvoiceRepository) GetInvoicesByBranch(branchID string, limit int) {}
