package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type InvTranSerialDetail struct {
	bun.BaseModel          `bun:"table:inv_tran_serial_detail"`
	CompanyId              string    `bun:"company_id,type:varchar(8)"`                 // Unique code that identifies a company.
	LocationId             float64   `bun:"location_id,type:decimal(19,0)"`             // Where was the used material located?
	SerialNumber           string    `bun:"serial_number,type:varchar(40)"`             // Serial number.
	TransactionNumber      int32     `bun:"transaction_number,type:int,default:((-1))"` // Transaction number from inv_tran. Not used after CC v7.0.
	DateCreated            time.Time `bun:"date_created,type:datetime"`                 // Indicates the date/time this record was created.
	DateLastModified       time.Time `bun:"date_last_modified,type:datetime"`           // Indicates the date/time this record was last modified.
	LastMaintainedBy       string    `bun:"last_maintained_by,type:varchar(30)"`        // ID of the user who last maintained this record
	DimensionTrackingKey   *int32    `bun:"dimension_tracking_key,type:int"`            // What -  if any -  dimension is used to track this serial number
	InvTranSerialDetailUid int32     `bun:"inv_tran_serial_detail_uid,type:int,pk"`     // Internal unique identifier for an inv_tran_serial_detail row.
	RowStatus              int16     `bun:"row_status,type:tinyint"`                    // Indicates current record status.
	DocumentLineSerialUid  int32     `bun:"document_line_serial_uid,type:int"`          // Unique identifier for the document_line_serial record
	InvMastUid             int32     `bun:"inv_mast_uid,type:int"`                      // Unique identifier for the item id.
	LotUid                 *int32    `bun:"lot_uid,type:int"`                           // lot uid when serial lot integration is used
}
