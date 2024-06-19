package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type ArReceiptsDetail struct {
	bun.BaseModel            `bun:"table:ar_receipts_detail"`
	ReceiptNumber            float64   `bun:"receipt_number,type:decimal(19,0),pk"`
	CustomerId               float64   `bun:"customer_id,type:decimal(19,0)"`                            // Customer paying invoice -  remitter
	InvoiceNo                string    `bun:"invoice_no,type:varchar(10),pk"`                            // Invoice being paid
	CompanyId                string    `bun:"company_id,type:varchar(8)"`                                // Unique code that identifies a company.
	PaymentAmount            float64   `bun:"payment_amount,type:decimal(19,4)"`                         // Amount of payment
	TermsAmount              float64   `bun:"terms_amount,type:decimal(19,4)"`                           // This is the terms amount
	AllowedAmount            float64   `bun:"allowed_amount,type:decimal(19,4)"`                         // Allowed amount accepted
	DeleteFlag               string    `bun:"delete_flag,type:char(1)"`                                  // Indicates whether this record is logically deleted
	DateCreated              time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified         time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy         string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	CurrencyVarianceAmtHome  float64   `bun:"currency_variance_amt_home,type:decimal(19,2),default:(0)"` // Stores variance amt in home currency from exchange rate fluctuation from invoice time to payment time
	CreatedBy                string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	TaxTermsTakenAmt         float64   `bun:"tax_terms_taken_amt,type:decimal(19,9),nullzero"`          // Terms for Taxes
	OtherChargeTermsTakenAmt float64   `bun:"other_charge_terms_taken_amt,type:decimal(19,9),nullzero"` // Terms for Other Charge Items
	TaxAmountPaid            float64   `bun:"tax_amount_paid,type:decimal(19,2),default:((0))"`         // column to hold how much of tax amount was paid during a payment.
}
