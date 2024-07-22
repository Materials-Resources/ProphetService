package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type ReversePaymentLine struct {
	bun.BaseModel         `bun:"table:reverse_payment_line"`
	ReversePaymentLineUid int32     `bun:"reverse_payment_line_uid,type:int,pk"`                         // Unique ID for Reverse Payment Line table
	ReversePaymentHdrUid  int32     `bun:"reverse_payment_hdr_uid,type:int"`                             // Unique ID for Reverse Payment Header from reverse_payment_hdr table
	InvoiceNo             string    `bun:"invoice_no,type:varchar(10)"`                                  // The Invoice for which Cash Receipt is being reversed
	ReverseAmt            float64   `bun:"reverse_amt,type:decimal(19,9),default:(0)"`                   // Amount that is being reversed
	ShipToId              *float64  `bun:"ship_to_id,type:decimal(19,0)"`                                // ShipToID for which the payment is reversed
	DepositNo             *string   `bun:"deposit_no,type:varchar(40)"`                                  // Deposit Number against which the amount is reversed
	ReversalTypeFlag      string    `bun:"reversal_type_flag,type:char(1),default:('R')"`                // Type of Cash Receipt Reversal; expected values: R-Reversed, P-Reapplied, C-Charged
	ReceiptNo             *float64  `bun:"receipt_no,type:decimal(19,0)"`                                // Cash Receipts Receipt No that is being reversed
	PaymentNo             *float64  `bun:"payment_no,type:decimal(19,0)"`                                // Cash Receipts Payment No that is being reversed
	DateCreated           time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy             string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified      time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy      string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
