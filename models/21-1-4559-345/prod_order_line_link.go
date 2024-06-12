package model

import (
	"github.com/uptrace/bun"
	"time"
)

type ProdOrderLineLink struct {
	bun.BaseModel       `bun:"table:prod_order_line_link"`
	ProdOrderNumber     float64   `bun:"prod_order_number,type:decimal(19,0),pk"`
	ProdOrderLineNumber float64   `bun:"prod_order_line_number,type:decimal(19,0),pk"`
	TransType           string    `bun:"trans_type,type:char,pk"`
	Quantity            float64   `bun:"quantity,type:decimal(19,9)"`
	DateCreated         time.Time `bun:"date_created,type:datetime"`
	DateLastModified    time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy    string    `bun:"last_maintained_by,type:varchar(30)"`
	RowStatusFlag       string    `bun:"row_status_flag,type:char"`
	TransactionUid      int32     `bun:"transaction_uid,type:int,pk"`
}
