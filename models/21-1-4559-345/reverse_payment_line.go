package model

import (
	"github.com/uptrace/bun"
	"time"
)

type ReversePaymentLine struct {
	bun.BaseModel         `bun:"table:reverse_payment_line"`
	ReversePaymentLineUid int32     `bun:"reverse_payment_line_uid,type:int,pk"`
	ReversePaymentHdrUid  int32     `bun:"reverse_payment_hdr_uid,type:int"`
	InvoiceNo             string    `bun:"invoice_no,type:varchar(10)"`
	ReverseAmt            float64   `bun:"reverse_amt,type:decimal(19,9),default:(0)"`
	ShipToId              float64   `bun:"ship_to_id,type:decimal(19,0),nullzero"`
	DepositNo             string    `bun:"deposit_no,type:varchar(40),nullzero"`
	ReversalTypeFlag      string    `bun:"reversal_type_flag,type:char,default:('R')"`
	ReceiptNo             float64   `bun:"receipt_no,type:decimal(19,0),nullzero"`
	PaymentNo             float64   `bun:"payment_no,type:decimal(19,0),nullzero"`
	DateCreated           time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy             string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified      time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy      string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
