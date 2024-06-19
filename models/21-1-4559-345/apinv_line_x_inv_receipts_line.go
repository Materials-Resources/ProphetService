package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type ApinvLineXInvReceiptsLine struct {
	bun.BaseModel           `bun:"table:apinv_line_x_inv_receipts_line"`
	ApinvLineXInvRecLineUid int32     `bun:"apinv_line_x_inv_rec_line_uid,type:int,autoincrement,identity,pk"` // Unique identifier for the record
	ReceiptNumber           float64   `bun:"receipt_number,type:decimal(19,0)"`                                // Inventory receipt number that this record relates to.
	ReceiptLine             int32     `bun:"receipt_line,type:int"`                                            // Line on the inventory receipt that this record relates to.
	ApinvLineUid            int32     `bun:"apinv_line_uid,type:int,nullzero"`                                 // Unique identifier of the voucher line that this record relates to.
	DateCreated             time.Time `bun:"date_created,type:datetime,default:(getdate())"`                   // Date and time the record was originally created
	CreatedBy               string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`             // User who created the record
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`             // Date and time the record was modified
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`     // User who last changed the record
}
