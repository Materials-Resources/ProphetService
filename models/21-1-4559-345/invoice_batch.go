package model

import (
	"github.com/uptrace/bun"
	"time"
)

type InvoiceBatch struct {
	bun.BaseModel                  `bun:"table:invoice_batch"`
	InvoiceBatchUid                int32     `bun:"invoice_batch_uid,type:int,pk"`
	InvoiceBatchNumber             int32     `bun:"invoice_batch_number,type:int"`
	InvoiceBatchDesc               string    `bun:"invoice_batch_desc,type:varchar(40)"`
	RowStatusFlag                  int32     `bun:"row_status_flag,type:int,default:(704)"`
	DateCreated                    time.Time `bun:"date_created,type:datetime"`
	DateLastModified               time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy               string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	ConsolidateByOrderFlag         string    `bun:"consolidate_by_order_flag,type:char,default:('N')"`
	ConsolidateCompletedOrdersFlag string    `bun:"consolidate_completed_orders_flag,type:char,default:('Y')"`
}
