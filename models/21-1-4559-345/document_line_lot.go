package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type DocumentLineLot struct {
	bun.BaseModel      `bun:"table:document_line_lot"`
	DocumentNo         int32     `bun:"document_no,type:int"`                            // Document number this lot quantity is associated with
	LineNo             int32     `bun:"line_no,type:int"`                                // Document line this lot quantity is associated with
	DocumentType       string    `bun:"document_type,type:char(2)"`                      // What is the type of the document (e.g. OR = order, IA = inventory adjustment)
	LotCd              string    `bun:"lot_cd,type:varchar(40)"`                         // What is the lot code for this document line lot?
	QtyToChange        float64   `bun:"qty_to_change,type:decimal(19,9)"`                // Quantity to change for this row (unit size = 1)
	QtyApplied         float64   `bun:"qty_applied,type:decimal(19,9)"`                  // Quantity that has been applied/completed for this row (unit size = 1)
	UnitQuantity       float64   `bun:"unit_quantity,type:decimal(19,9)"`                // Lot quantity in terms of unit size
	UnitOfMeasure      string    `bun:"unit_of_measure,type:varchar(8)"`                 // What is the unit of measure for this row?
	DocumentCd         string    `bun:"document_cd,type:varchar(10),nullzero"`           // String ID of the document
	DateCreated        time.Time `bun:"date_created,type:datetime"`                      // Indicates the date/time this record was created.
	DateLastModified   time.Time `bun:"date_last_modified,type:datetime"`                // Indicates the date/time this record was last modified.
	LastMaintainedBy   string    `bun:"last_maintained_by,type:varchar(30)"`             // ID of the user who last maintained this record
	DocumentLineLotUid int32     `bun:"document_line_lot_uid,type:int,pk"`               // What is the unique identifier for this document line lot?
	UnitSize           float64   `bun:"unit_size,type:decimal(19,4)"`                    // SKU size for the unit of measure
	SubLineNo          int32     `bun:"sub_line_no,type:int"`                            // Row for detail records for the main row (assembly components, etc)
	SequenceNo         int32     `bun:"sequence_no,type:int,default:(0)"`                // User defined sort order of the stages.
	LotUid             int32     `bun:"lot_uid,type:int,nullzero"`                       // Uniquer row ID for the lot table.  Used as a staging area for unapproved transactions.
	QtyFromTags        float64   `bun:"qty_from_tags,type:decimal(19,9),default:((-1))"` // Stores the qty that was posted from a tag posting.  Used to tell the difference between manual and tag postings.
	ItemRevisionUid    int32     `bun:"item_revision_uid,type:int,nullzero"`             // Column holds item revision uid of lot.
	CurrentbinUid      int32     `bun:"currentbin_uid,type:int,nullzero"`                // Current Bin UID
	QaStatus           string    `bun:"qa_status,type:varchar(255),nullzero"`            // Determines the status of the QA process for this lot - typically for a sales order.
	SubInvoiceNo       string    `bun:"sub_invoice_no,type:varchar(40),nullzero"`        // Sub Invoice number for Pro Forma Invoice
}
