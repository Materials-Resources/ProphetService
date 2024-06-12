package model

import (
	"github.com/uptrace/bun"
	"time"
)

type OeLineRma struct {
	bun.BaseModel      `bun:"table:oe_line_rma"`
	OeLineRmaUid       int32     `bun:"oe_line_rma_uid,type:int,pk"`
	OeLineUid          int32     `bun:"oe_line_uid,type:int"`
	RmaLinkedOeLineUid int32     `bun:"rma_linked_oe_line_uid,type:int"`
	DateCreated        time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	DateLastModified   time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy   string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	RowStatusFlag      int32     `bun:"row_status_flag,type:int,default:('704')"`
	InvoiceLineUid     int32     `bun:"invoice_line_uid,type:int,nullzero"`
	RetrievedByWms     string    `bun:"retrieved_by_wms,type:char,default:('N')"`
}
