package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type DocumentLineSerial struct {
	bun.BaseModel         `bun:"table:document_line_serial"`
	DocumentNo            float64   `bun:"document_no,type:decimal(19,0)"`                  // Document number this serial number is associated with
	LineNo                int32     `bun:"line_no,type:int"`                                // Document line this serial number is associated with
	DocumentType          string    `bun:"document_type,type:char(2)"`                      // What is the type of the document (e.g. OR = order, IA = inventory adjustment)
	SerialNumber          string    `bun:"serial_number,type:varchar(40)"`                  // Serial Number associated with this document
	DocumentCd            *string   `bun:"document_cd,type:varchar(10)"`                    // String ID of the document
	DateCreated           time.Time `bun:"date_created,type:datetime"`                      // Indicates the date/time this record was created.
	DateLastModified      time.Time `bun:"date_last_modified,type:datetime"`                // Indicates the date/time this record was last modified.
	LastMaintainedBy      string    `bun:"last_maintained_by,type:varchar(30)"`             // ID of the user who last maintained this record
	DimensionTrackingKey  *int32    `bun:"dimension_tracking_key,type:int"`                 // What -  if any -  dimension is used to track this seri
	DocumentLineSerialUid int32     `bun:"document_line_serial_uid,type:int,pk"`            // Unique identifier for this document line serial
	RowStatus             int16     `bun:"row_status,type:tinyint"`                         // Indicates current record status.
	SubLineNo             int32     `bun:"sub_line_no,type:int"`                            // Row for detail records for the main row (assembly components, etc)
	SequenceNo            int32     `bun:"sequence_no,type:int,default:(0)"`                // What sequence should the process be performed in -  for this stage?
	SerialNumberUid       *int32    `bun:"serial_number_uid,type:int"`                      // Uniquer row ID for the serial table.  Used as a staging area for unapproved transactions.
	QtyFromTags           float64   `bun:"qty_from_tags,type:decimal(19,9),default:((-1))"` // Stores the qty that was posted from a tag posting.  Used to tell the difference between manual and tag postings.
	LotUid                *int32    `bun:"lot_uid,type:int"`                                // lot uid when serial lot integration is used
	RfBinUid              *int32    `bun:"rf_bin_uid,type:int"`                             // Indicates the rf bin that picked the serial number.
	DepositBinUid         *int32    `bun:"deposit_bin_uid,type:int"`                        // Keep the Slab Bin ID until the PO receipt is approved
}
