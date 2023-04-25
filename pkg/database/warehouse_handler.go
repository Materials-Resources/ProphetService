package database

import (
	"context"
	"fmt"
	"github.com/uptrace/bun"
)

type WarehouseHandler struct {
	db *bun.DB
}

type SelectPickTicketResults struct {
	OrderId string
}

type SelectReceivingReportResults struct {
	Products []*SelectReceivingReportResultsProduct
}

type SelectReceivingReportResultsProduct struct {
	Id  int32
	Qty float32
	Uom string
}

func NewWarehouseHandler(db *bun.DB) *WarehouseHandler {
	return &WarehouseHandler{db: db}
}

func (h *WarehouseHandler) SelectPickTicket(ctx context.Context, id int32) (*PickTicket, error) {
	model := new(PickTicket)
	err := h.db.NewSelect().Model(model).Where("pick_ticket_no = ?", id).Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not query Pick Ticket: %v", err)
	}
	return model, nil
}

func (h *WarehouseHandler) SelectReceivingReport(ctx context.Context, id int32) (*InventoryReceipt, error) {
	model := new(InventoryReceipt)

	err := h.db.NewSelect().Model(model).Where("receipt_number = ?", id).Relation("Items").Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not query Receiving Report: %v", err)
	}

	results := &SelectReceivingReportResults{}

	for _, item := range model.Items {
		results.Products = append(results.Products, &SelectReceivingReportResultsProduct{
			Id:  item.InvMastUid,
			Qty: item.Qty,
			Uom: item.Uom,
		})
	}
	return model, nil

}

func (h *WarehouseHandler) SelectReceivingReportAllocatedOrders(ctx context.Context, receiptId int32, productId int32, primaryBins []string) (*InventoryReceiptLine, error) {
	model := new(InventoryReceiptLine)

	query := h.db.NewSelect().
		Model(model).
		Relation("OrderTransactions").
		Where("inventory_receipt_line.receipt_number = ? AND inventory_receipt_line.inv_mast_uid = ?", receiptId, productId)

	if len(primaryBins) > 0 {
		query.Relation("InvMast.InvLoc").Where("inv_mast__inv_loc.primary_bin in (?)", bun.In(primaryBins))
	}

	if err := query.Scan(ctx); err != nil {
		return nil, err
	}

	return model, nil
}
