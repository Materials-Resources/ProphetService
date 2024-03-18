package repository

import (
	"context"
	"errors"
	"strconv"

	"github.com/materials-resources/s_prophet/infrastructure/db/prophet_19_1_3668"
	"github.com/materials-resources/s_prophet/internal/inventory"
	"github.com/materials-resources/s_prophet/internal/inventory/domain"
	"github.com/uptrace/bun"
)

type bunInventoryRepository struct {
	db *bun.DB
}

func (b bunInventoryRepository) SelectProductStockById(ctx context.Context, id []int32) ([]*domain.ProductStock,
	error,
) {
	var mInvLoc []invLoc
	err := b.db.NewSelect().Model(&mInvLoc).Column("inv_mast_uid",
		"qty_on_hand",
	).Where("inv_mast_uid IN (?)", bun.In(id)).Scan(ctx)
	if err != nil {
		return nil, err
	}
	var dProductStock []*domain.ProductStock
	for _, invLoc := range mInvLoc {
		dProductStock = append(dProductStock, invLoc.WriteToProductStock())

	}
	return dProductStock, nil
}

func NewBunInventoryRepository(db *bun.DB) inventory.Repository {
	return &bunInventoryRepository{db: db}
}

func (b bunInventoryRepository) FindReceiptById(id string) (*domain.ValidatedInventoryReceipt, error) {
	ctx := context.Background()
	dbInventoryReceiptsHdr := new(prophet_19_1_3668.InventoryReceiptsHdr)

	dbID, err := strconv.ParseFloat(id, 64)
	if err != nil {
		return nil, errors.New("ID provided is invalid")
	}
	err = b.db.NewSelect().Model(dbInventoryReceiptsHdr).Column("receipt_number").
		Relation("InventoryReceiptsLineItems",
			func(q *bun.SelectQuery) *bun.SelectQuery {
				return q.Column("receipt_number", "inv_mast_uid", "qty_received", "unit_of_measure")
			},
		).
		Relation("InventoryReceiptsLineItems.Ira", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Column("document_no",
				"sub_document_no",
				"trans_type",
				"inv_mast_uid",
			)
		},
		).
		Relation("InventoryReceiptsLineItems.InvLoc", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Column("inv_mast_uid", "primary_bin")
		},
		).
		Relation("InventoryReceiptsLineItems.InvLoc.InvMast", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Column("item_id", "item_desc")
		},
		).
		Where("receipt_number = ?",
			dbID,
		).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return FromDBReceipt(dbInventoryReceiptsHdr)
}
