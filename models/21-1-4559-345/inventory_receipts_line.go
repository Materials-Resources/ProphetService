package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type InventoryReceiptsLine struct {
	bun.BaseModel             `bun:"table:inventory_receipts_line"`
	ReceiptNumber             float64   `bun:"receipt_number,type:decimal(19,0),pk"`
	LineNumber                int32     `bun:"line_number,type:int,pk"`                                   // The line number on the order.
	QtyReceived               float64   `bun:"qty_received,type:decimal(19,9)"`                           // Quantity received
	UnitOfMeasure             string    `bun:"unit_of_measure,type:varchar(8)"`                           // What is the unit of measure for this row?
	UnitSize                  float64   `bun:"unit_size,type:decimal(19,4)"`                              // Size of the quantity unit of measure
	UnitQuantity              float64   `bun:"unit_quantity,type:decimal(19,9)"`                          // Quantity received in the purchasing UOM.
	DateCreated               time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified          time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy          string    `bun:"last_maintained_by,type:varchar(30)"`                       // ID of the user who last maintained this record
	VouchComplete             string    `bun:"vouch_complete,type:char(1)"`                               // Set to Y when the line item is fully converted to voucher.
	QtyVouched                float64   `bun:"qty_vouched,type:decimal(19,9)"`                            // Quantity converted to voucher so far.
	PoLineNumber              int32     `bun:"po_line_number,type:int"`                                   // The line item from the originating purchase order.
	UnitCost                  float64   `bun:"unit_cost,type:decimal(19,9),nullzero"`                     // Cost per unit.
	ExtendedCost              float64   `bun:"extended_cost,type:decimal(19,4),nullzero"`                 // The cost for a line item, calculated by multiplying the quantity ordered, requested, or shipped by the item’s Unit Cost.
	FreightAmount             float64   `bun:"freight_amount,type:decimal(19,4),nullzero"`                // The portion of the overall freight amount prorated to this line item.
	AmountVouched             float64   `bun:"amount_vouched,type:decimal(19,4),nullzero"`                // The amount previously vouched through the Post Rec
	FreightAmountVouched      float64   `bun:"freight_amount_vouched,type:decimal(19,4),nullzero"`        // The portion of the freight amount that has been converted to voucher.
	UnitCostDisplay           float64   `bun:"unit_cost_display,type:decimal(19,9),nullzero"`             // Holds the display unit price in the currency of the receipt.
	ExtendedCostDisplay       float64   `bun:"extended_cost_display,type:decimal(19,4),nullzero"`         // Holds the display extended cost.  It is necessary
	FreightAmountDisplay      float64   `bun:"freight_amount_display,type:decimal(19,4),nullzero"`        // Holds the display freight amount. It is necessary
	InvMastUid                int32     `bun:"inv_mast_uid,type:int"`                                     // Unique identifier for the item id.
	QtyDisassembled           float64   `bun:"qty_disassembled,type:decimal(19,9),default:(0)"`           // Amount of quantity that is disassembled.
	QtyLost                   float64   `bun:"qty_lost,type:decimal(19,9),nullzero"`                      // store the Qty Lost per receipt for the MSP Receiving Report
	MspLandedCost             float64   `bun:"msp_landed_cost,type:decimal(19,9),nullzero"`               // The landed cost for a receipt when moving material from a po stage.
	SubstituteItem            string    `bun:"substitute_item,type:char(1),default:('N')"`                // Has item been substituted for another item?
	OriginalInvMastUid        int32     `bun:"original_inv_mast_uid,type:int,nullzero"`                   // Original item unique identifier
	PricingUnit               string    `bun:"pricing_unit,type:varchar(8)"`                              // Maintains the po_line pricing unit of measure when material is received.
	PricingUnitSize           float64   `bun:"pricing_unit_size,type:decimal(19,4)"`                      // Maintains the po_line pricing unit size when material is received.
	SubstituteLineNumber      int32     `bun:"substitute_line_number,type:int,nullzero"`                  // to link the substitute line with the original line number
	DateReceiptRemoved        time.Time `bun:"date_receipt_removed,type:datetime,nullzero"`               // The date a line was removed from receipt.
	RetrievedByWms            string    `bun:"retrieved_by_wms,type:char(1),default:('N')"`               // column to indicate if the production order component was retrieved by wms
	TransferFlag              string    `bun:"transfer_flag,type:char(1),nullzero"`                       // A flag to indicate whether a transfer is needed or not.
	CountryOfOrigin           string    `bun:"country_of_origin,type:varchar(255),nullzero"`              // Custom column to the country of origin for the item in PO receipts and Transfer receipts.
	VesselLandedCostPostedAmt float64   `bun:"vessel_landed_cost_posted_amt,type:decimal(19,4),nullzero"` // When this receipt line was from a container receipt, this column will store the vessel receipt level landed cost applied to this particular receipt.
	WrongPartNoFlag           string    `bun:"wrong_part_no_flag,type:char(1),default:('N')"`             // Wrong Part Number indicator for PO Received lines.
	PalletId                  string    `bun:"pallet_id,type:varchar(30),nullzero"`                       // Value to be entered PO Receipts window in WWMS
	CreatedBy                 string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	LabelQty                  int32     `bun:"label_qty,type:int,nullzero"` // Custom: Indicates the number of labels to print for corresponding line item.
	InventoryReceiptsLineUid  int32     `bun:"inventory_receipts_line_uid,type:int,autoincrement,scanonly"`
	ReceiptRemovedBy          string    `bun:"receipt_removed_by,type:varchar(255),nullzero"`            // Indicates user id that reversed this inventory receipt line (if it has been reversed).
	DirectShipFreightAmount   float64   `bun:"direct_ship_freight_amount,type:decimal(19,4),nullzero"`   // Custom: Indicates the freight added to receipt from linked oe line item.
	ExcludeFromLandedCostFlag string    `bun:"exclude_from_landed_cost_flag,type:char(1),default:('N')"` // Exclude line from landed cost calculation
}
