package model

import (
	"github.com/uptrace/bun"
	"time"
)

type InventoryReceiptsLine struct {
	bun.BaseModel             `bun:"table:inventory_receipts_line"`
	ReceiptNumber             float64   `bun:"receipt_number,type:decimal(19,0),pk"`
	LineNumber                int32     `bun:"line_number,type:int,pk"`
	QtyReceived               float64   `bun:"qty_received,type:decimal(19,9)"`
	UnitOfMeasure             string    `bun:"unit_of_measure,type:varchar(8)"`
	UnitSize                  float64   `bun:"unit_size,type:decimal(19,4)"`
	UnitQuantity              float64   `bun:"unit_quantity,type:decimal(19,9)"`
	DateCreated               time.Time `bun:"date_created,type:datetime"`
	DateLastModified          time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy          string    `bun:"last_maintained_by,type:varchar(30)"`
	VouchComplete             string    `bun:"vouch_complete,type:char"`
	QtyVouched                float64   `bun:"qty_vouched,type:decimal(19,9)"`
	PoLineNumber              int32     `bun:"po_line_number,type:int"`
	UnitCost                  float64   `bun:"unit_cost,type:decimal(19,9),nullzero"`
	ExtendedCost              float64   `bun:"extended_cost,type:decimal(19,4),nullzero"`
	FreightAmount             float64   `bun:"freight_amount,type:decimal(19,4),nullzero"`
	AmountVouched             float64   `bun:"amount_vouched,type:decimal(19,4),nullzero"`
	FreightAmountVouched      float64   `bun:"freight_amount_vouched,type:decimal(19,4),nullzero"`
	UnitCostDisplay           float64   `bun:"unit_cost_display,type:decimal(19,9),nullzero"`
	ExtendedCostDisplay       float64   `bun:"extended_cost_display,type:decimal(19,4),nullzero"`
	FreightAmountDisplay      float64   `bun:"freight_amount_display,type:decimal(19,4),nullzero"`
	InvMastUid                int32     `bun:"inv_mast_uid,type:int"`
	QtyDisassembled           float64   `bun:"qty_disassembled,type:decimal(19,9),default:(0)"`
	QtyLost                   float64   `bun:"qty_lost,type:decimal(19,9),nullzero"`
	MspLandedCost             float64   `bun:"msp_landed_cost,type:decimal(19,9),nullzero"`
	SubstituteItem            string    `bun:"substitute_item,type:char,default:('N')"`
	OriginalInvMastUid        int32     `bun:"original_inv_mast_uid,type:int,nullzero"`
	PricingUnit               string    `bun:"pricing_unit,type:varchar(8)"`
	PricingUnitSize           float64   `bun:"pricing_unit_size,type:decimal(19,4)"`
	SubstituteLineNumber      int32     `bun:"substitute_line_number,type:int,nullzero"`
	DateReceiptRemoved        time.Time `bun:"date_receipt_removed,type:datetime,nullzero"`
	RetrievedByWms            string    `bun:"retrieved_by_wms,type:char,default:('N')"`
	TransferFlag              string    `bun:"transfer_flag,type:char,nullzero"`
	CountryOfOrigin           string    `bun:"country_of_origin,type:varchar(255),nullzero"`
	VesselLandedCostPostedAmt float64   `bun:"vessel_landed_cost_posted_amt,type:decimal(19,4),nullzero"`
	WrongPartNoFlag           string    `bun:"wrong_part_no_flag,type:char,default:('N')"`
	PalletId                  string    `bun:"pallet_id,type:varchar(30),nullzero"`
	CreatedBy                 string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	LabelQty                  int32     `bun:"label_qty,type:int,nullzero"`
	InventoryReceiptsLineUid  int32     `bun:"inventory_receipts_line_uid,type:int,identity"`
	ReceiptRemovedBy          string    `bun:"receipt_removed_by,type:varchar(255),nullzero"`
	DirectShipFreightAmount   float64   `bun:"direct_ship_freight_amount,type:decimal(19,4),nullzero"`
	ExcludeFromLandedCostFlag string    `bun:"exclude_from_landed_cost_flag,type:char,default:('N')"`
}
