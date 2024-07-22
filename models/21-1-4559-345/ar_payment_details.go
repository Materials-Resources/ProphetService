package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type ArPaymentDetails struct {
	bun.BaseModel               `bun:"table:ar_payment_details"`
	PaymentNumber               float64    `bun:"payment_number,type:decimal(19,0),pk"`                      // Unique identifier - system generated
	PaymentTypeId               float64    `bun:"payment_type_id,type:decimal(19,0)"`                        // System generated payment type identifier
	PaymentDesc                 string     `bun:"payment_desc,type:varchar(60)"`                             // A description field
	PaymentDate                 time.Time  `bun:"payment_date,type:datetime"`                                // Date payment accepted
	PaymentAmount               float64    `bun:"payment_amount,type:decimal(19,4)"`                         // Amount being paid to invoice
	TermsAmount                 *float64   `bun:"terms_amount,type:decimal(19,4)"`                           // Terms amount for invoice
	AllowedAmount               *float64   `bun:"allowed_amount,type:decimal(19,4)"`                         // Allowed amount for invoice
	CheckNumber                 *string    `bun:"check_number,type:varchar(30)"`                             // Check number used to pay
	CcName                      *string    `bun:"cc_name,type:varchar(30)"`                                  // Name of credit card used
	CcNumber                    *string    `bun:"cc_number,type:varchar(30)"`                                // Credit Card Number
	CcExpirationDate            *time.Time `bun:"cc_expiration_date,type:datetime"`                          // Credit Card Expiration Date
	CcAuthorizedDate            *time.Time `bun:"cc_authorized_date,type:datetime"`                          // Credit Card Authorization Date
	CcAuthorizedNumber          *string    `bun:"cc_authorized_number,type:varchar(30)"`                     // Credt Card Authorization Number
	Period                      *float64   `bun:"period,type:decimal(3,0)"`                                  // In which period did the inventory transfer occur?
	YearForPeriod               *float64   `bun:"year_for_period,type:decimal(4,0)"`                         // What year does the period belong to?
	DepositNumber               *string    `bun:"deposit_number,type:varchar(40)"`                           // Deposit number used when depositing money to bank
	DeleteFlag                  string     `bun:"delete_flag,type:char(1)"`                                  // Indicates whether this record is logically deleted
	DateCreated                 time.Time  `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified            time.Time  `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy            string     `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	MemoAmount                  float64    `bun:"memo_amount,type:decimal(19,4)"`                            // Enter the memo amount
	PreprocessedPaymentAmount   *float64   `bun:"preprocessed_payment_amount,type:decimal(19,4)"`            // If user specifies payment amount in RMA Entry window, this column will hold the payment amount for un-invoiced RMA if payment type is credit card.
	CurrencyLineUid             *int32     `bun:"currency_line_uid,type:int"`                                // Used to identify currency info for the payment, namely the exchange rate for a foreign currency transaction.
	CurrencyVarianceAmtHome     float64    `bun:"currency_variance_amt_home,type:decimal(19,2),default:(0)"` // Stores variance amt in home currency from exchange rate fluctuation from invoice time to payment time
	CreatedBy                   *string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	SourceTypeCd                *int32     `bun:"source_type_cd,type:int"`                           // To track down how a record got created. Cash Receipts / Imports / Remittances etc
	CustomerPaymentExchRate     *float64   `bun:"customer_payment_exch_rate,type:decimal(19,9)"`     // Holds exchange rate b/w customer and payment currencies.
	PaymentSourceCd             *int32     `bun:"payment_source_cd,type:int"`                        // Indicates who made the payment.  In most cases, it is the customer (so column is null) but can be the Homeowner.
	StoreCreditNumber           *int32     `bun:"store_credit_number,type:int"`                      // Unique identifier of Store Credit Table record associated with the Payment.
	ChangeAmount                *float64   `bun:"change_amount,type:decimal(19,2)"`                  // Amount returned as changes from new cash pymt in shipping
	PaymentAccountId            *string    `bun:"payment_account_id,type:varchar(255)"`              // Identifies a account associated with the payment (typically from an external source, such as Element Payment Services)
	WebCcProcessorUid           *int32     `bun:"web_cc_processor_uid,type:int"`                     // Credit card processor ID used on the website for transaction
	TaxTermsAmount              float64    `bun:"tax_terms_amount,type:decimal(19,2),default:((0))"` // column to hold the total terms on tax amount that was taken during cash receipts
	TaxAmountPaid               float64    `bun:"tax_amount_paid,type:decimal(19,2),default:((0))"`  // column to hold how much of tax amount was paid during a payment.
	PayAmtOnlyFlag              *string    `bun:"pay_amt_only_flag,type:char(1)"`                    // Flag to allow the user to specify that only the amount entered into the Payment Amount field will be charged to the credit card.
	RoundedAmount               *float64   `bun:"rounded_amount,type:decimal(19,4)"`                 // Rounded value to apply to payment amount to calculate actual payment when price rounding rules is enabled.
	BilltrustPaymentTypeCd      *int32     `bun:"billtrust_payment_type_cd,type:int"`                // Custom: Indicates the type of bill trust payment received.
	MultiCurrencyId             *int32     `bun:"multi_currency_id,type:int"`                        // Currency ID the payment was taken in.
	MultiCurrencyPaymentAmount  *float64   `bun:"multi_currency_payment_amount,type:decimal(19,9)"`  // Payment Amount based on the the multi currency ID
	MultiCurrencyExchangeRate   *float64   `bun:"multi_currency_exchange_rate,type:decimal(28,17)"`  // Exchange rate of the payment's currency to the customer's currency.
	MultiCurrencyVarianceAmount *float64   `bun:"multi_currency_variance_amount,type:decimal(19,9)"` // Variance Amount between the customer amount to home and the multi currency payment amount to home.
	ExchangeRateOverrideDate    *time.Time `bun:"exchange_rate_override_date,type:datetime"`         // Indicate that the user choose to override the effective exchange rate, based on this date, applied to a foreign currency payment.  Currency_line_uid should reflect the exchange rate for this date.
	CcReturnTransactionId       *string    `bun:"cc_return_transaction_id,type:varchar(255)"`        // The transaction ID of a previous credit card payment being used to return against.
	CcRefundParentPaymentNo     *float64   `bun:"cc_refund_parent_payment_no,type:decimal(19,0)"`    // For credit card refund records, link to the associated original credit card payment record.
	EwingCouponUid              *int32     `bun:"ewing_coupon_uid,type:int"`                         // FK to new ewing_coupon table
}
