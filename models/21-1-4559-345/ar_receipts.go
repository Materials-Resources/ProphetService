package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type ArReceipts struct {
	bun.BaseModel      `bun:"table:ar_receipts"`
	ReceiptNumber      float64   `bun:"receipt_number,type:decimal(19,0),pk"`                      // System-generated number that identifies a receipt against a purchase order, transfer, cash receipt, or a confirmation against a direct ship purchase order.
	PaymentNumber      float64   `bun:"payment_number,type:decimal(19,0)"`                         // Payment number for this receipt
	CompanyId          string    `bun:"company_id,type:varchar(8)"`                                // Unique code that identifies a company.
	RemitterId         float64   `bun:"remitter_id,type:decimal(19,0)"`                            // Company paying
	DateReceived       time.Time `bun:"date_received,type:datetime"`                               // Date money received
	BankNo             *float64  `bun:"bank_no,type:decimal(19,0)"`                                // Enter a valid bank number
	GlBankAccountNo    *string   `bun:"gl_bank_account_no,type:varchar(32)"`                       // GL bank account number
	ArAccountNo        *string   `bun:"ar_account_no,type:varchar(32)"`                            // What accouncts receivable account is this invoice to be credited to ?
	TermsAccountNo     *string   `bun:"terms_account_no,type:varchar(32)"`                         // The default account number to post terms discount for a new customer
	AllowedAccountNo   *string   `bun:"allowed_account_no,type:varchar(32)"`                       // Allowed account number
	Approved           string    `bun:"approved,type:char(1)"`                                     // Has this product order been approved?
	DeleteFlag         string    `bun:"delete_flag,type:char(1)"`                                  // Indicates whether this record is logically deleted
	DateCreated        time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified   time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy   string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	CreatedBy          *string   `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	ReversePaymentFlag *string   `bun:"reverse_payment_flag,type:char(1),default:('N')"` // Indicates whether a Cash Receipt has been reversed / check returned or not
	ClearedBank        string    `bun:"cleared_bank,type:char(1),default:('N')"`         // Has this payment been cleared by the bank?
	ClearedPeriod      *int16    `bun:"cleared_period,type:smallint"`                    // The period in which the payment cleared the bank.
	ClearedYear        *int16    `bun:"cleared_year,type:smallint"`                      // The year in which the payment cleared the bank.
}
