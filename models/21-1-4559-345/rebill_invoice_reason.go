package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type RebillInvoiceReason struct {
	bun.BaseModel          `bun:"table:rebill_invoice_reason"`
	RebillInvoiceReasonUid int32     `bun:"rebill_invoice_reason_uid,type:int,autoincrement,pk"`          // Unique identifier for this row.
	CompanyId              string    `bun:"company_id,type:varchar(8),unique"`                            // The company for which this reason code applies.
	ReasonCode             string    `bun:"reason_code,type:varchar(255),unique"`                         // The user-defined Reason ID.
	ReasonDesc             string    `bun:"reason_desc,type:varchar(255)"`                                // The description of the reason.
	RowStatus              int16     `bun:"row_status,type:smallint,default:((704))"`                     // Indicates the status of the reason code.
	DateCreated            time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy              string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified       time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy       string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
