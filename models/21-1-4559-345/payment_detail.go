package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type PaymentDetail struct {
	bun.BaseModel           `bun:"table:payment_detail"`
	CompanyNo               string    `bun:"company_no,type:varchar(8)"`                                // Unique code that identifies a company.
	BankNo                  float64   `bun:"bank_no,type:decimal(19,0)"`                                // Default bank number
	CheckNo                 string    `bun:"check_no,type:varchar(255)"`                                // This is the check number
	InvoiceNo               string    `bun:"invoice_no,type:varchar(32)"`                               // Invoice number that is associated with payment
	PoNo                    string    `bun:"po_no,type:varchar(50),nullzero"`                           // What is the purchase order for this location(s)?
	DateCreated             time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	AmountPaid              float64   `bun:"amount_paid,type:decimal(19,4),nullzero"`                   // Payment amount in terms of the company currency ID
	VoucherNo               string    `bun:"voucher_no,type:varchar(10),nullzero"`                      // Voucher number that is associated with payment
	HomeAmtPaid             float64   `bun:"home_amt_paid,type:decimal(19,4),nullzero"`                 // Payment amount in terms of the company currency ID
	TermsAmountTaken        float64   `bun:"terms_amount_taken,type:decimal(19,4),nullzero"`            // Terms amount in terms of the company currency ID
	AmountPaidDisplay       float64   `bun:"amount_paid_display,type:decimal(19,4),nullzero"`           // Payment amount in terms of the vendor currency ID
	TermsAmountTakenDisplay float64   `bun:"terms_amount_taken_display,type:decimal(19,4),nullzero"`    // Terms amount in terms of the vendor currency ID
	PaymentDetailUid        int32     `bun:"payment_detail_uid,type:int,pk"`                            // Unique Identifier for the table.
	CurrencyVarianceAmount  float64   `bun:"currency_variance_amount,type:decimal(19,4)"`               // Variance amount due to the foreign currency exchange rate
}
