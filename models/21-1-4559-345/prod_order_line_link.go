package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type ProdOrderLineLink struct {
	bun.BaseModel       `bun:"table:prod_order_line_link"`
	ProdOrderNumber     float64   `bun:"prod_order_number,type:decimal(19,0),pk"`      // Production Order number to be linked
	ProdOrderLineNumber float64   `bun:"prod_order_line_number,type:decimal(19,0),pk"` // Production Order line number to be linked
	TransType           string    `bun:"trans_type,type:char(1),pk"`                   // Transaction Type of document this Prod Order is linked to (Order/Production Order/etc)
	Quantity            float64   `bun:"quantity,type:decimal(19,9)"`                  // Quantity to be linked
	DateCreated         time.Time `bun:"date_created,type:datetime"`                   // Indicates the date/time this record was created.
	DateLastModified    time.Time `bun:"date_last_modified,type:datetime"`             // Indicates the date/time this record was last modified.
	LastMaintainedBy    string    `bun:"last_maintained_by,type:varchar(30)"`          // ID of the user who last maintained this record
	RowStatusFlag       string    `bun:"row_status_flag,type:char(1)"`                 // Indicates current record status.
	TransactionUid      int32     `bun:"transaction_uid,type:int,pk"`                  // Unique identifier for the transaction this Prod Order is linked to
}
