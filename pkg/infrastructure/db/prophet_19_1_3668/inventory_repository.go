package prophet_19_1_3668

import (
	"context"
	"errors"
	"strconv"

	"github.com/materials-resources/s_prophet/pkg/domain/entities"
	"github.com/materials-resources/s_prophet/pkg/domain/repositories"
	"github.com/materials-resources/s_prophet/pkg/infrastructure/db/prophet_19_1_3668/model"
	"github.com/uptrace/bun"
)

type bunInventoryRepository struct {
	db *bun.DB
}

func NewBunInventoryRepository(db *bun.DB) repositories.InventoryRepository {
	return &bunInventoryRepository{db: db}
}

func (b bunInventoryRepository) FindReceiptById(id string) (*entities.ValidatedInventoryReceipt, error) {
	ctx := context.Background()
	dbInventoryReceiptsHdr := new(model.InventoryReceiptsHdr)

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
