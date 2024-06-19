package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type CustomerCallInvDetail struct {
	bun.BaseModel            `bun:"table:customer_call_inv_detail"`
	CustomerCallInvDetailUid int32     `bun:"customer_call_inv_detail_uid,type:int,pk"` // Unique Identifier of customer_call_inv_detail table.
	CustomerCallUid          int32     `bun:"customer_call_uid,type:int"`               // Unique Identifier of customer_call table.
	InvoiceNo                string    `bun:"invoice_no,type:varchar(10)"`              // Invoice number for customer call.
	DateCreated              time.Time `bun:"date_created,type:datetime"`               // Indicates the date/time this record was created.
	DateLastModified         time.Time `bun:"date_last_modified,type:datetime"`         // Indicates the date/time this record was last modified.
	LastMaintainedBy         string    `bun:"last_maintained_by,type:varchar(30)"`      // ID of the user who last maintained this record
}
