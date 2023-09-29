package repository

import (
	"context"
	"fmt"

	"github.com/materials-resources/s_prophet/pkg/internal/core/port/repository"
	"github.com/materials-resources/s_prophet/pkg/internal/infra/repository/model"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mssqldialect"
)

type receivingRepository struct {
	bun *bun.DB
	db  repository.Database
}

func (r receivingRepository) ReadReceivingReceipt() error {
	inventoryReceiptsHdr := new(model.InventoryReceiptsHdr)
	ctx := context.Background()
	err := r.bun.NewSelect().Model(inventoryReceiptsHdr).Relation("InventoryReceiptsLineItems").Relation(
		"InventoryReceiptsLineItems.InvMast",
	).Where(
		"receipt_number = ?",
		5000011,
	).Scan(ctx)
	if err != nil {
		fmt.Print(err)
		return err
	}
	fmt.Println(inventoryReceiptsHdr.ReceiptNumber)
	for _, line := range inventoryReceiptsHdr.InventoryReceiptsLineItems {
		fmt.Printf(
			"Line: %d, ID: %d, Item Desc: %s \n",
			line.LineNumber,
			line.InvMastUid,
			line.InvMast.ItemDesc,
		)
	}
	return nil
}

func NewReceivingRepository(db repository.Database) repository.ReceivingRepository {
	bundb := bun.NewDB(
		db.GetDB(),
		mssqldialect.New(),
	)
	return &receivingRepository{
		db:  db,
		bun: bundb,
	}
}
