package model

import (
	"github.com/uptrace/bun"
	"time"
)

type ArPaymentDetails struct {
	bun.BaseModel               `bun:"table:ar_payment_details"`
	PaymentNumber               float64   `bun:"payment_number,type:decimal(19,0),pk"`
	PaymentTypeId               float64   `bun:"payment_type_id,type:decimal(19,0)"`
	PaymentDesc                 string    `bun:"payment_desc,type:varchar(60)"`
	PaymentDate                 time.Time `bun:"payment_date,type:datetime"`
	PaymentAmount               float64   `bun:"payment_amount,type:decimal(19,4)"`
	TermsAmount                 float64   `bun:"terms_amount,type:decimal(19,4),nullzero"`
	AllowedAmount               float64   `bun:"allowed_amount,type:decimal(19,4),nullzero"`
	CheckNumber                 string    `bun:"check_number,type:varchar(30),nullzero"`
	CcName                      string    `bun:"cc_name,type:varchar(30),nullzero"`
	CcNumber                    string    `bun:"cc_number,type:varchar(30),nullzero"`
	CcExpirationDate            time.Time `bun:"cc_expiration_date,type:datetime,nullzero"`
	CcAuthorizedDate            time.Time `bun:"cc_authorized_date,type:datetime,nullzero"`
	CcAuthorizedNumber          string    `bun:"cc_authorized_number,type:varchar(30),nullzero"`
	Period                      float64   `bun:"period,type:decimal(3,0),nullzero"`
	YearForPeriod               float64   `bun:"year_for_period,type:decimal(4,0),nullzero"`
	DepositNumber               string    `bun:"deposit_number,type:varchar(40),nullzero"`
	DeleteFlag                  string    `bun:"delete_flag,type:char"`
	DateCreated                 time.Time `bun:"date_created,type:datetime"`
	DateLastModified            time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy            string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	MemoAmount                  float64   `bun:"memo_amount,type:decimal(19,4)"`
	PreprocessedPaymentAmount   float64   `bun:"preprocessed_payment_amount,type:decimal(19,4),nullzero"`
	CurrencyLineUid             int32     `bun:"currency_line_uid,type:int,nullzero"`
	CurrencyVarianceAmtHome     float64   `bun:"currency_variance_amt_home,type:decimal(19,2),default:(0)"`
	CreatedBy                   string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	SourceTypeCd                int32     `bun:"source_type_cd,type:int,nullzero"`
	CustomerPaymentExchRate     float64   `bun:"customer_payment_exch_rate,type:decimal(19,9),nullzero"`
	PaymentSourceCd             int32     `bun:"payment_source_cd,type:int,nullzero"`
	StoreCreditNumber           int32     `bun:"store_credit_number,type:int,nullzero"`
	ChangeAmount                float64   `bun:"change_amount,type:decimal(19,2),nullzero"`
	PaymentAccountId            string    `bun:"payment_account_id,type:varchar(255),nullzero"`
	WebCcProcessorUid           int32     `bun:"web_cc_processor_uid,type:int,nullzero"`
	TaxTermsAmount              float64   `bun:"tax_terms_amount,type:decimal(19,2),default:((0))"`
	TaxAmountPaid               float64   `bun:"tax_amount_paid,type:decimal(19,2),default:((0))"`
	PayAmtOnlyFlag              string    `bun:"pay_amt_only_flag,type:char,nullzero"`
	RoundedAmount               float64   `bun:"rounded_amount,type:decimal(19,4),nullzero"`
	BilltrustPaymentTypeCd      int32     `bun:"billtrust_payment_type_cd,type:int,nullzero"`
	MultiCurrencyId             int32     `bun:"multi_currency_id,type:int,nullzero"`
	MultiCurrencyPaymentAmount  float64   `bun:"multi_currency_payment_amount,type:decimal(19,9),nullzero"`
	MultiCurrencyExchangeRate   float64   `bun:"multi_currency_exchange_rate,type:decimal(28,17),nullzero"`
	MultiCurrencyVarianceAmount float64   `bun:"multi_currency_variance_amount,type:decimal(19,9),nullzero"`
	ExchangeRateOverrideDate    time.Time `bun:"exchange_rate_override_date,type:datetime,nullzero"`
	CcReturnTransactionId       string    `bun:"cc_return_transaction_id,type:varchar(255),nullzero"`
	CcRefundParentPaymentNo     float64   `bun:"cc_refund_parent_payment_no,type:decimal(19,0),nullzero"`
	EwingCouponUid              int32     `bun:"ewing_coupon_uid,type:int,nullzero"`
}
