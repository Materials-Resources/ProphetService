package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type OeLineRma struct {
	bun.BaseModel      `bun:"table:oe_line_rma"`
	OeLineRmaUid       int32     `bun:"oe_line_rma_uid,type:int,pk"`                               // Unique ID for this oe_line_rma record.
	OeLineUid          int32     `bun:"oe_line_uid,type:int"`                                      // Internal unique value to identify an order line.
	RmaLinkedOeLineUid int32     `bun:"rma_linked_oe_line_uid,type:int"`                           // Unique ID for the oe_line associated with this oe_line_rma record.
	DateCreated        time.Time `bun:"date_created,type:datetime,default:(getdate())"`            // Indicates the date/time this record was created.
	DateLastModified   time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`      // Indicates the date/time this record was last modified.
	LastMaintainedBy   string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	RowStatusFlag      int32     `bun:"row_status_flag,type:int,default:('704')"`                  // Indicates current record status.
	InvoiceLineUid     int32     `bun:"invoice_line_uid,type:int,nullzero"`                        // Foreign Key to invoice_line: Column will provide more efficient processing for said feature
	RetrievedByWms     string    `bun:"retrieved_by_wms,type:char(1),default:('N')"`               // Indicates if the production order component was retrieved by wms
}
