package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type Payments struct {
	bun.BaseModel              `bun:"table:payments"`
	CompanyNo                  string    `bun:"company_no,type:varchar(8),pk"`                                   // Unique code that identifies a company.
	BankNo                     float64   `bun:"bank_no,type:decimal(19,0),pk"`                                   // Bank Number.
	CheckNo                    string    `bun:"check_no,type:varchar(255),pk"`                                   // What was the check number?
	CheckDate                  time.Time `bun:"check_date,type:datetime"`                                        // What is the date of the payment -  as marked on the
	Void                       string    `bun:"void,type:char(1)"`                                               // Has this payment been voided?
	CheckAmount                float64   `bun:"check_amount,type:decimal(19,4)"`                                 // Check amount in terms of the company currency ID
	ClearedBank                string    `bun:"cleared_bank,type:char(1)"`                                       // Has this payment been cleared by the bank?
	TermsTakenAmount           float64   `bun:"terms_taken_amount,type:decimal(19,4)"`                           // Terms amount in terms of the company currency ID
	VendorId                   float64   `bun:"vendor_id,type:decimal(19,0)"`                                    // What is the unique vendor identifier for this row?
	ApAccountNo                string    `bun:"ap_account_no,type:varchar(32)"`                                  // Enter an AP account number
	DateCreated                time.Time `bun:"date_created,type:datetime"`                                      // Indicates the date/time this record was created.
	DateLastModified           time.Time `bun:"date_last_modified,type:datetime"`                                // Indicates the date/time this record was last modified.
	LastMaintainedBy           string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"`   // ID of the user who last maintained this record
	Period                     float64   `bun:"period,type:decimal(3,0),nullzero"`                               // Which period does this quota apply to?
	YearForPeriod              float64   `bun:"year_for_period,type:decimal(4,0),nullzero"`                      // What year does the period belong to?
	CurrencyId                 float64   `bun:"currency_id,type:decimal(19,0),nullzero"`                         // What is the unique currency identifier for this ro
	HomeChkAmount              float64   `bun:"home_chk_amount,type:decimal(19,4),nullzero"`                     // Check amount in terms of the company currency ID
	VoidPeriod                 float64   `bun:"void_period,type:decimal(3,0),nullzero"`                          // What period was this payment voided in?
	VoidYear                   float64   `bun:"void_year,type:decimal(4,0),nullzero"`                            // What yar was this payment voided in?
	StatementClearedDate       time.Time `bun:"statement_cleared_date,type:datetime,nullzero"`                   // Date in which the check was cleared
	TermsTakenAmountDisplay    float64   `bun:"terms_taken_amount_display,type:decimal(19,4)"`                   // Terms amount in terms of the vendor currency ID
	CheckAmountDisplay         float64   `bun:"check_amount_display,type:decimal(19,4)"`                         // Check amount in terms of the vendor currency ID
	CurrencyVarianceAmount     float64   `bun:"currency_variance_amount,type:decimal(19,4)"`                     // Variance amount due to the foreign currency exchange rate
	TransmissionMethod         int32     `bun:"transmission_method,type:int,nullzero"`                           // column is used to identify payments that have been sent using EDI
	Reason                     string    `bun:"reason,type:varchar(255),nullzero"`                               // to hold the reason for a voided check
	ClearedPeriod              int16     `bun:"cleared_period,type:smallint,nullzero"`                           // The period in which the payment cleared the bank.
	ClearedYear                int16     `bun:"cleared_year,type:smallint,nullzero"`                             // The year in which the payment cleared the bank.
	CheckMailedDate            time.Time `bun:"check_mailed_date,type:datetime,nullzero"`                        // Custom Feature 28424: Date check was mailed to vendor
	HomeCurrencyAmount         float64   `bun:"home_currency_amount,type:decimal(19,9),nullzero"`                // Custom column to store the total home currency amount for a vendor check
	CalculatedExchangeRate     float64   `bun:"calculated_exchange_rate,type:decimal(19,9),nullzero"`            // Custom exchange rate calculated using home currency amount and vendor check total
	IvaWithheldAmount          float64   `bun:"iva_withheld_amount,type:decimal(19,4),default:((0))"`            // IVA Withheld Amount
	IvaWithheldAmountDisplay   float64   `bun:"iva_withheld_amount_display,type:decimal(19,4),default:((0))"`    // IVA Withheld Amount (Foreign Currency)
	IvaTermsAmountTaken        float64   `bun:"iva_terms_amount_taken,type:decimal(19,4),default:((0))"`         // IVA Terms Amount
	IvaTermsAmountTakenDisplay float64   `bun:"iva_terms_amount_taken_display,type:decimal(19,4),default:((0))"` // IVA Terms Amount (Foreign Currency)
	VoidDate                   time.Time `bun:"void_date,type:datetime,nullzero"`                                // Date Check was Voided
	BankGlAcctNo               string    `bun:"bank_gl_acct_no,type:varchar(255),nullzero"`                      // Bank GL acct debited at time of payment.
	ExchangeRate               float64   `bun:"exchange_rate,type:decimal(19,6),nullzero"`                       // Stores the exchange rate used.
	SourceTypeCd               int32     `bun:"source_type_cd,type:int,nullzero"`                                // Code (from code_p21 table) which indicates where the Check was created
}
