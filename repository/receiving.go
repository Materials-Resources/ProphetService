package repository

import (
	"context"
	"fmt"
	"strconv"

	"github.com/materials-resources/s_prophet/core/domain"
	"github.com/materials-resources/s_prophet/core/port/repository"
	"github.com/materials-resources/s_prophet/model"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mssqldialect"
	"github.com/uptrace/bun/extra/bundebug"
)

type receivingRepository struct {
	bun *bun.DB
	db  repository.Database
}

func (r receivingRepository) ReadReceipt(ctx context.Context, id string) (*domain.ReceivingReceipt, error) {
	inventoryReceiptsHdr := new(model.InventoryReceiptsHdr)
	err := r.bun.NewSelect().Model(inventoryReceiptsHdr).
		Relation("InventoryReceiptsLineItems").
		Relation(
			"InventoryReceiptsLineItems.InvMast",
		).
		Relation("InventoryReceiptsLineItems.InvMast.InvLoc").
		Relation("InventoryReceiptsLineItems.Ira").
		Relation("InventoryReceiptsLineItems.Ira.OeHdr").
		Where(
			"inventory_receipts_hdr.receipt_number = ?",
			5000011,
		).Scan(ctx)
	if err != nil {
		fmt.Print(err)
		return nil, err
	}
	return inventoryReceiptsHdrToDomain(*inventoryReceiptsHdr), nil
}

func NewReceivingRepository(db repository.Database) repository.ReceivingRepository {
	bundb := bun.NewDB(
		db.GetDB(),
		mssqldialect.New(),
	)
	bundb.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))

	return &receivingRepository{
		db:  db,
		bun: bundb,
	}
}

func inventoryReceiptsHdrToDomain(inventoryReceiptsHdr model.InventoryReceiptsHdr) *domain.ReceivingReceipt {
	allocatedOrders := func(orders []*model.InvTran) []*domain.ReceivingReceiptAllocatedOrder {
		var res []*domain.ReceivingReceiptAllocatedOrder
		for _, order := range orders {
			res = append(
				res,
				&domain.ReceivingReceiptAllocatedOrder{
					Id: strconv.FormatFloat(
						order.DocumentNo,
						'f',
						-1,
						64,
					),
					UnitsAllocated: order.QtyAllocated,
					UnitLabel:      order.UnitOfMeasure,
					Name:           order.OeHdr.Ship2Name.String,
				},
			)
		}
		return res
	}
	receivingReceipt := new(domain.ReceivingReceipt)

	receivingReceipt.Id = strconv.FormatFloat(
		inventoryReceiptsHdr.ReceiptNumber,
		'f',
		-1,
		64,
	)
	for _, item := range inventoryReceiptsHdr.InventoryReceiptsLineItems {
		receivingReceipt.Products = append(
			receivingReceipt.Products,
			&domain.ReceivingReceiptProduct{
				Name:            item.InvMast.ItemDesc,
				Sn:              item.InvMast.ItemId,
				Id:              strconv.Itoa(int(item.InvMastUid)),
				UnitsReceived:   item.UnitQuantity,
				UnitLabel:       item.UnitOfMeasure,
				PrimaryBin:      item.InvMast.InvLoc.PrimaryBin.String,
				AllocatedOrders: allocatedOrders(item.Ira),
			},
		)
	}
	return receivingReceipt
}
