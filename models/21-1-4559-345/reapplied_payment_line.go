package model

import (
	"github.com/uptrace/bun"
	"time"
)

type ReappliedPaymentLine struct {
	bun.BaseModel           `bun:"table:reapplied_payment_line"`
	ReappliedPaymentLineUid int32     `bun:"reapplied_payment_line_uid,type:int,pk"`
	ReversePaymentHdrUid    int32     `bun:"reverse_payment_hdr_uid,type:int"`
	InvoiceNo               string    `bun:"invoice_no,type:varchar(10),nullzero"`
	ShipToId                float64   `bun:"ship_to_id,type:decimal(19,0),nullzero"`
	BranchId                string    `bun:"branch_id,type:varchar(8),nullzero"`
	Amt                     float64   `bun:"amt,type:decimal(19,9),default:(0)"`
	TermsAmt                float64   `bun:"terms_amt,type:decimal(19,9),nullzero"`
	AllowedAmt              float64   `bun:"allowed_amt,type:decimal(19,9),nullzero"`
	CreditStatus            string    `bun:"credit_status,type:varchar(20),nullzero"`
	DateCreated             time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy               string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	ReappliedReceiptNo      float64   `bun:"reapplied_receipt_no,type:decimal(19,0),nullzero"`
	ReappliedPaymentNo      float64   `bun:"reapplied_payment_no,type:decimal(19,0),nullzero"`
}
