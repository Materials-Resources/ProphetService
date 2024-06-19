package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type DocumentLineBin struct {
	bun.BaseModel         `bun:"table:document_line_bin"`
	DocumentNo            float64   `bun:"document_no,type:decimal(19,0)"`                       // Document number this bin quantity is associated with
	LineNo                int32     `bun:"line_no,type:int"`                                     // Document line this bin quantity is associated with
	DocumentType          string    `bun:"document_type,type:char(2)"`                           // What is the type of the document (e.g. OR = order, IA = inventory adjustment)
	BinCd                 string    `bun:"bin_cd,type:varchar(40)"`                              // Bin code
	QtyApplied            float64   `bun:"qty_applied,type:decimal(19,9)"`                       // Quantity that has been applied/completed for this row (unit size = 1)
	UnitQuantity          float64   `bun:"unit_quantity,type:decimal(19,9)"`                     // Bin quantity in terms of unit size
	UnitOfMeasure         string    `bun:"unit_of_measure,type:varchar(8)"`                      // What is the unit of measure for this row?
	DocumentCd            string    `bun:"document_cd,type:varchar(10),nullzero"`                // String ID of the document
	DateCreated           time.Time `bun:"date_created,type:datetime"`                           // Indicates the date/time this record was created.
	DateLastModified      time.Time `bun:"date_last_modified,type:datetime"`                     // Indicates the date/time this record was last modified.
	LastMaintainedBy      string    `bun:"last_maintained_by,type:varchar(30)"`                  // ID of the user who last maintained this record
	DocumentLineBinUid    int32     `bun:"document_line_bin_uid,type:int,pk"`                    // What is the unique -  internal identifier for this document line bin?
	UnitSize              float64   `bun:"unit_size,type:decimal(19,4)"`                         // SKU size for the unit of measure
	QtyToChange           float64   `bun:"qty_to_change,type:decimal(19,9)"`                     // Quantity to change for this row (unit size = 1)
	SubLineNo             int32     `bun:"sub_line_no,type:int"`                                 // Row for detail records for the main row (assembly components, etc)
	RfQtyPicked           float64   `bun:"rf_qty_picked,type:decimal(19,9),default:(0)"`         // Quantity that has been picked on an RF device
	QtyFromTags           float64   `bun:"qty_from_tags,type:decimal(19,9),default:((-1))"`      // Stores the qty that was posted from a tag posting.  Used to tell the difference between manual and tag postings.
	SourceDlbUid          int32     `bun:"source_dlb_uid,type:int,nullzero"`                     // UID of document_line_bin of source location. Used e.g. when in transfer receipts: if this is a bin that was received, then the source_dlb_uid refers to the associated bin that was shipped.
	CreatedBy             string    `bun:"created_by,type:varchar(255),default:(suser_sname())"` // User who created the record
	PrintedFlag           string    `bun:"printed_flag,type:char(1),nullzero"`                   // Custom column to indicate if the record has been printed in Transfer window
	WorkOrderUid          int32     `bun:"work_order_uid,type:int,nullzero"`                     // Work order UID for custom picking process
	PickStatus            string    `bun:"pick_status,type:varchar(2),nullzero"`                 // Pick status for custom picking process
	AssignedWorkstationId string    `bun:"assigned_workstation_id,type:varchar(255),nullzero"`   // Workstation ID for custom picking process
}
