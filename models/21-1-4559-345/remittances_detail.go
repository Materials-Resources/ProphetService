package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type RemittancesDetail struct {
	bun.BaseModel           `bun:"table:remittances_detail"`
	RemittanceNumber        float64   `bun:"remittance_number,type:decimal(19,0),pk"`                       // What remittance does this remittance detail belong to?
	InvoiceNo               string    `bun:"invoice_no,type:varchar(10),pk"`                                // Invoice being paid
	PaymentAmount           float64   `bun:"payment_amount,type:decimal(19,4)"`                             // Amount of payment
	TermsAmount             float64   `bun:"terms_amount,type:decimal(19,4)"`                               // Terms amount for invoice
	AllowedAmount           float64   `bun:"allowed_amount,type:decimal(19,4)"`                             // allowed amount for invoice
	DeleteFlag              string    `bun:"delete_flag,type:char(1)"`                                      // Indicates whether this record is logically deleted
	DateCreated             time.Time `bun:"date_created,type:datetime"`                                    // Indicates the date/time this record was created.
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime"`                              // Indicates the date/time this record was last modified.
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"` // ID of the user who last maintained this record
	CurrencyVarianceAmtHome float64   `bun:"currency_variance_amt_home,type:decimal(19,2),default:(0)"`     // Stores variance amt in home currency from exchange rate fluctuation from invoice time to payment time
	TaxAmountPaid           float64   `bun:"tax_amount_paid,type:decimal(19,2),default:((0))"`              // column to hold how much of tax amount was paid during a payment.
}
