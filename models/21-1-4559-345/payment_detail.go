package model

import (
	"github.com/uptrace/bun"
	"time"
)

type PaymentDetail struct {
	bun.BaseModel           `bun:"table:payment_detail"`
	CompanyNo               string    `bun:"company_no,type:varchar(8)"`
	BankNo                  float64   `bun:"bank_no,type:decimal(19,0)"`
	CheckNo                 string    `bun:"check_no,type:varchar(255)"`
	InvoiceNo               string    `bun:"invoice_no,type:varchar(32)"`
	PoNo                    string    `bun:"po_no,type:varchar(50),nullzero"`
	DateCreated             time.Time `bun:"date_created,type:datetime"`
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	AmountPaid              float64   `bun:"amount_paid,type:decimal(19,4),nullzero"`
	VoucherNo               string    `bun:"voucher_no,type:varchar(10),nullzero"`
	HomeAmtPaid             float64   `bun:"home_amt_paid,type:decimal(19,4),nullzero"`
	TermsAmountTaken        float64   `bun:"terms_amount_taken,type:decimal(19,4),nullzero"`
	AmountPaidDisplay       float64   `bun:"amount_paid_display,type:decimal(19,4),nullzero"`
	TermsAmountTakenDisplay float64   `bun:"terms_amount_taken_display,type:decimal(19,4),nullzero"`
	PaymentDetailUid        int32     `bun:"payment_detail_uid,type:int,pk"`
	CurrencyVarianceAmount  float64   `bun:"currency_variance_amount,type:decimal(19,4)"`
}
