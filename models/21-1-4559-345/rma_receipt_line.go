package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type RmaReceiptLine struct {
	bun.BaseModel      `bun:"table:rma_receipt_line"`
	RmaReceiptLineUid  int32     `bun:"rma_receipt_line_uid,type:int,pk"`                             // Unique internal ID number
	RmaReceiptHdrUid   int32     `bun:"rma_receipt_hdr_uid,type:int"`                                 // FK to rma_receipt_hdr.rma_receipt_hdr_uid - link to associated rma_receipt_hdr record
	OeLineUid          int32     `bun:"oe_line_uid,type:int"`                                         // FK to oe_line.oe_line_uid - link to associated RMA line record
	ReceiptLineNo      int32     `bun:"receipt_line_no,type:int"`                                     // Receipt line number.  May not be the same as oe_line.line_no.
	QtyReceived        float64   `bun:"qty_received,type:decimal(19,9),default:(0)"`                  // Receipt quantity
	QtyToStock         float64   `bun:"qty_to_stock,type:decimal(19,9),default:(0)"`                  // Amount of receipt qty to return-to-stock.
	QtyToAdjust        float64   `bun:"qty_to_adjust,type:decimal(19,9),default:(0)"`                 // Amount of receipt qty not returned-to-stock.
	QtyToSupplier      float64   `bun:"qty_to_supplier,type:decimal(19,9),default:(0)"`               // Amount of receipt qty to return to supplier via and inventory return
	FreightIn          float64   `bun:"freight_in,type:decimal(19,9),default:(0)"`                    // In freight associated w/receipt line.
	ReasonAdjustmentId string    `bun:"reason_adjustment_id,type:varchar(5),nullzero"`                // Reason ID associated w/qty_to_adjust.
	DateCreated        time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy          string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified   time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy   string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	RowStatusFlag      int32     `bun:"row_status_flag,type:int,default:(704)"`                       // Specifies record active status.
	ParentOeLineUid    int32     `bun:"parent_oe_line_uid,type:int,default:(null)"`                   // For assembly component lines, link to associated parent rma_receipt_line record.
	UnitOfMeasure      string    `bun:"unit_of_measure,type:varchar(8),nullzero"`                     // Unit of measure chosen to describe the measurement of qtys on this receipt line.
	UnitSize           float64   `bun:"unit_size,type:decimal(19,9),nullzero"`                        // The number of SKUs associated w/the unit of measure defined for this receipt line.
	QtyToTransfer      float64   `bun:"qty_to_transfer,type:decimal(19,9),nullzero"`                  // Custom column to account amount of receipt qty to transfer to the destination location
	CoreItemCost       float64   `bun:"core_item_cost,type:decimal(19,9),nullzero"`                   // Custom column for core item supplier cost.
	CountryOfOrigin    string    `bun:"country_of_origin,type:varchar(255),nullzero"`                 // Custom column to store country of origin for item in RMA receipts.
}
