package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type InventoryReturnLine struct {
	bun.BaseModel          `bun:"table:inventory_return_line"`
	InventoryReturnLineUid int32     `bun:"inventory_return_line_uid,type:int,pk"`                     // Unique identifier for inventory_return_line
	InventoryReturnHdrUid  int32     `bun:"inventory_return_hdr_uid,type:int,unique"`                  // Used to link to the inventory_return_header record
	LineNumber             int32     `bun:"line_number,type:int,unique"`                               // The line number on the order.
	InvMastUid             int32     `bun:"inv_mast_uid,type:int"`                                     // Unique identifier for the item id.
	ExtendedDescription    string    `bun:"extended_description,type:varchar(255),nullzero"`           // Extended description for the item on the return.
	QuantityUom            string    `bun:"quantity_uom,type:varchar(8)"`                              // The quantity unit of measure
	QuantityUnitSize       float64   `bun:"quantity_unit_size,type:decimal(19,4)"`                     // The quantity unit size
	UnitQuantity           float64   `bun:"unit_quantity,type:decimal(19,9),default:(0)"`              // The quantity in units on the return
	QtyToReturn            float64   `bun:"qty_to_return,type:decimal(19,9),default:(0)"`              // The quantity in skus on the return
	QtyPicked              float64   `bun:"qty_picked,type:decimal(19,9),default:(0)"`                 // Quantity Picked
	QtyReturned            float64   `bun:"qty_returned,type:decimal(19,9),default:(0)"`               // The quantity in skus returned to the vendor.
	QtyVouched             float64   `bun:"qty_vouched,type:decimal(19,9),default:(0)"`                // The quantity in skus that was vouched.
	UnitPrice              float64   `bun:"unit_price,type:decimal(19,9),default:(0)"`                 // What is the unit price for this line item?
	UnitPriceDisplay       float64   `bun:"unit_price_display,type:decimal(19,9),default:(0)"`         // Holds the values that are displayed on the window.
	PriceUom               string    `bun:"price_uom,type:varchar(8)"`                                 // Price unit of measure
	PriceUnitSize          float64   `bun:"price_unit_size,type:decimal(19,4)"`                        // Price unit size
	ExtendedPrice          float64   `bun:"extended_price,type:decimal(19,2),default:(0)"`             // The extended price for the line
	PriceEdited            string    `bun:"price_edited,type:char(1),default:('N')"`                   // Indicates whether the price was edited
	ReturnDate             time.Time `bun:"return_date,type:datetime,default:(getdate())"`             // Date line was entered.
	ExpectedDate           time.Time `bun:"expected_date,type:datetime,default:(getdate())"`           // What is the expected date that this relatinship between a process and a transaction should be completed?
	SupplierPartNumber     string    `bun:"supplier_part_number,type:varchar(255),nullzero"`           // The part number according to the supplier.
	RowStatusFlag          int32     `bun:"row_status_flag,type:int,default:(702)"`                    // Indicates current record status.
	DateCreated            time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified       time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy       string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	RetrievedByWms         string    `bun:"retrieved_by_wms,type:char(1),default:('N')"`               // column to indicate if the production order component was retrieved by wms
	CoreItemCost           float64   `bun:"core_item_cost,type:decimal(19,9),nullzero"`                // Custom column for core item supplier cost.
}
