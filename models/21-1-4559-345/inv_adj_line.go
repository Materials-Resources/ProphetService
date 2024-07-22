package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type InvAdjLine struct {
	bun.BaseModel          `bun:"table:inv_adj_line"`
	AdjustmentNumber       float64   `bun:"adjustment_number,type:decimal(19,0),pk"`                    // System generated ID for the adjustment.
	LineNumber             int32     `bun:"line_number,type:int,pk"`                                    // System generated line number, with adjustment number, unique ID for this row.
	UnitQuantity           float64   `bun:"unit_quantity,type:decimal(19,9)"`                           // Adjustment Amount.  In UOMs
	Cost                   float64   `bun:"cost,type:decimal(19,9)"`                                    // Unit cost
	UnitOfMeasure          string    `bun:"unit_of_measure,type:varchar(8)"`                            // SKU for this line
	DateCreated            time.Time `bun:"date_created,type:datetime"`                                 // Indicates the date/time this record was created.
	DateLastModified       time.Time `bun:"date_last_modified,type:datetime"`                           // Indicates the date/time this record was last modified.
	LastMaintainedBy       string    `bun:"last_maintained_by,type:varchar(30)"`                        // ID of the user who last maintained this record
	Quantity               float64   `bun:"quantity,type:decimal(19,9)"`                                // Quantity in SKUs
	InvMastUid             int32     `bun:"inv_mast_uid,type:int"`                                      // Unique identifier for the item id.
	RowStatusFlag          int32     `bun:"row_status_flag,type:int,default:(701)"`                     // Indicates whether the quantity on an unapproved line has been edited or not for the RF inventory counting window.
	QtyCountedFlag         *string   `bun:"qty_counted_flag,type:char(1)"`                              // Used by custom to indicate that the qty on an adjustment line created during physical count was counted.
	VendorId               *float64  `bun:"vendor_id,type:decimal(19,0)"`                               // This is a custom column to specify vendor id for item qty to adjust in RMA receipt window
	DeallocateQtyFlag      string    `bun:"deallocate_qty_flag,type:char(1),default:('N')"`             // Determines whether to deallocate quantity before processing the adjustment for this line.
	WwmsDeallocateFailFlag string    `bun:"wwms_deallocate_fail_flag,type:char(1),default:('N')"`       // Determines whether the deallocation routine failed for a wireless transaction.
	ParentTransLineNo      *int32    `bun:"parent_trans_line_no,type:int"`                              // Line Number of Associated Transaction
	ActualPhysicalCountQty *float64  `bun:"actual_physical_count_qty,type:decimal(19,9),default:((0))"` // Custom column to store the actual count
	RecountFlag            *string   `bun:"recount_flag,type:char(1),default:('N')"`
	TimesRecounted         *int32    `bun:"times_recounted,type:int,default:((0))"`            // Increment everytime an item is counted during recount on wwms.
	ReasonId               *string   `bun:"reason_id,type:varchar(5)"`                         // Custom (F65493): FK to reason.id.  Link to associated reason record.  For physical counts, the reason for this lines adjustment.
	Comment                *string   `bun:"comment,type:varchar(255)"`                         // Custom (F65493): user defined comment for this instance
	AdjustFoundItemFlag    string    `bun:"adjust_found_item_flag,type:char(1),default:('N')"` // For cycle/physical counts, determines if a found item will be adjusted.
}
