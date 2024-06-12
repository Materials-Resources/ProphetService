package model

import (
	"github.com/uptrace/bun"
	"time"
)

type Payments struct {
	bun.BaseModel              `bun:"table:payments"`
	CompanyNo                  string    `bun:"company_no,type:varchar(8),pk"`
	BankNo                     float64   `bun:"bank_no,type:decimal(19,0),pk"`
	CheckNo                    string    `bun:"check_no,type:varchar(255),pk"`
	CheckDate                  time.Time `bun:"check_date,type:datetime"`
	Void                       string    `bun:"void,type:char"`
	CheckAmount                float64   `bun:"check_amount,type:decimal(19,4)"`
	ClearedBank                string    `bun:"cleared_bank,type:char"`
	TermsTakenAmount           float64   `bun:"terms_taken_amount,type:decimal(19,4)"`
	VendorId                   float64   `bun:"vendor_id,type:decimal(19,0)"`
	ApAccountNo                string    `bun:"ap_account_no,type:varchar(32)"`
	DateCreated                time.Time `bun:"date_created,type:datetime"`
	DateLastModified           time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy           string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"`
	Period                     float64   `bun:"period,type:decimal(3,0),nullzero"`
	YearForPeriod              float64   `bun:"year_for_period,type:decimal(4,0),nullzero"`
	CurrencyId                 float64   `bun:"currency_id,type:decimal(19,0),nullzero"`
	HomeChkAmount              float64   `bun:"home_chk_amount,type:decimal(19,4),nullzero"`
	VoidPeriod                 float64   `bun:"void_period,type:decimal(3,0),nullzero"`
	VoidYear                   float64   `bun:"void_year,type:decimal(4,0),nullzero"`
	StatementClearedDate       time.Time `bun:"statement_cleared_date,type:datetime,nullzero"`
	TermsTakenAmountDisplay    float64   `bun:"terms_taken_amount_display,type:decimal(19,4)"`
	CheckAmountDisplay         float64   `bun:"check_amount_display,type:decimal(19,4)"`
	CurrencyVarianceAmount     float64   `bun:"currency_variance_amount,type:decimal(19,4)"`
	TransmissionMethod         int32     `bun:"transmission_method,type:int,nullzero"`
	Reason                     string    `bun:"reason,type:varchar(255),nullzero"`
	ClearedPeriod              int16     `bun:"cleared_period,type:smallint,nullzero"`
	ClearedYear                int16     `bun:"cleared_year,type:smallint,nullzero"`
	CheckMailedDate            time.Time `bun:"check_mailed_date,type:datetime,nullzero"`
	HomeCurrencyAmount         float64   `bun:"home_currency_amount,type:decimal(19,9),nullzero"`
	CalculatedExchangeRate     float64   `bun:"calculated_exchange_rate,type:decimal(19,9),nullzero"`
	IvaWithheldAmount          float64   `bun:"iva_withheld_amount,type:decimal(19,4),default:((0))"`
	IvaWithheldAmountDisplay   float64   `bun:"iva_withheld_amount_display,type:decimal(19,4),default:((0))"`
	IvaTermsAmountTaken        float64   `bun:"iva_terms_amount_taken,type:decimal(19,4),default:((0))"`
	IvaTermsAmountTakenDisplay float64   `bun:"iva_terms_amount_taken_display,type:decimal(19,4),default:((0))"`
	VoidDate                   time.Time `bun:"void_date,type:datetime,nullzero"`
	BankGlAcctNo               string    `bun:"bank_gl_acct_no,type:varchar(255),nullzero"`
	ExchangeRate               float64   `bun:"exchange_rate,type:decimal(19,6),nullzero"`
	SourceTypeCd               int32     `bun:"source_type_cd,type:int,nullzero"`
}
