package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type InvoiceBatch struct {
	bun.BaseModel                  `bun:"table:invoice_batch"`
	InvoiceBatchUid                int32     `bun:"invoice_batch_uid,type:int,pk"`                                // System generated unique identifier for invoice batches
	InvoiceBatchNumber             int32     `bun:"invoice_batch_number,type:int,unique"`                         // Code which the system uses to identify invoice batches
	InvoiceBatchDesc               string    `bun:"invoice_batch_desc,type:varchar(40)"`                          // Description of the the invoice batch
	RowStatusFlag                  int32     `bun:"row_status_flag,type:int,default:(704)"`                       // Indicates current record status.
	DateCreated                    time.Time `bun:"date_created,type:datetime"`                                   // Indicates the date/time this record was created.
	DateLastModified               time.Time `bun:"date_last_modified,type:datetime"`                             // Indicates the date/time this record was last modified.
	LastMaintainedBy               string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`    // ID of the user who last maintained this record
	ConsolidateByOrderFlag         string    `bun:"consolidate_by_order_flag,type:char(1),default:('N')"`         // Consolidate invoice by order
	ConsolidateCompletedOrdersFlag string    `bun:"consolidate_completed_orders_flag,type:char(1),default:('Y')"` // Consolidate Completed Orders Only
}
