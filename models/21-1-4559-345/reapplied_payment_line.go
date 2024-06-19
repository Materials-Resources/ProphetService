package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type ReappliedPaymentLine struct {
	bun.BaseModel           `bun:"table:reapplied_payment_line"`
	ReappliedPaymentLineUid int32     `bun:"reapplied_payment_line_uid,type:int,pk"`                       // Unique ID for Reapplied Payment Line table
	ReversePaymentHdrUid    int32     `bun:"reverse_payment_hdr_uid,type:int"`                             // Unique ID for Reverse Payment Header from reverse_payment_hdr table
	InvoiceNo               string    `bun:"invoice_no,type:varchar(10),nullzero"`                         // The Invoice for which Cash Receipt is being reapplied
	ShipToId                float64   `bun:"ship_to_id,type:decimal(19,0),nullzero"`                       // ShipToID for which the payment is reapplied
	BranchId                string    `bun:"branch_id,type:varchar(8),nullzero"`                           // BranchID for which check return fee is being charged; Check return only
	Amt                     float64   `bun:"amt,type:decimal(19,9),default:(0)"`                           // Reapplied amount or returned check fee amount
	TermsAmt                float64   `bun:"terms_amt,type:decimal(19,9),nullzero"`                        // Terms amount being reapplied
	AllowedAmt              float64   `bun:"allowed_amt,type:decimal(19,9),nullzero"`                      // Allowed amount being reapplied
	CreditStatus            string    `bun:"credit_status,type:varchar(20),nullzero"`                      // Modified credit status for ShipTo
	DateCreated             time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy               string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	ReappliedReceiptNo      float64   `bun:"reapplied_receipt_no,type:decimal(19,0),nullzero"`             // Cash Receipts Receipt No for the reapplied record
	ReappliedPaymentNo      float64   `bun:"reapplied_payment_no,type:decimal(19,0),nullzero"`             // Cash Receipts Payment No for the reapplied record
}
