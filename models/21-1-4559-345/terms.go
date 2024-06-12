package model

import (
	"github.com/uptrace/bun"
	"time"
)

type Terms struct {
	bun.BaseModel              `bun:"table:terms"`
	TermsId                    string    `bun:"terms_id,type:varchar(2),pk"`
	TermsDesc                  string    `bun:"terms_desc,type:varchar(20)"`
	DiscountPct                float64   `bun:"discount_pct,type:decimal(19,4),nullzero"`
	DiscountDays               float64   `bun:"discount_days,type:decimal(3,0),nullzero"`
	NetDays                    float64   `bun:"net_days,type:decimal(3,0),nullzero"`
	DeleteFlag                 string    `bun:"delete_flag,type:char"`
	DateCreated                time.Time `bun:"date_created,type:datetime"`
	DateLastModified           time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy           string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"`
	DayOfMonth                 float64   `bun:"day_of_month,type:decimal(2,0),nullzero"`
	Months                     float64   `bun:"months,type:decimal(2,0),nullzero"`
	TermsDayOfMonth            float64   `bun:"terms_day_of_month,type:decimal(2,0),nullzero"`
	TermsMonths                float64   `bun:"terms_months,type:decimal(2,0),nullzero"`
	NetDayOfMonth2             float64   `bun:"net_day_of_month2,type:decimal(2,0),nullzero"`
	TermsDayOfMonth2           float64   `bun:"terms_day_of_month2,type:decimal(2,0),nullzero"`
	DownpaymentPct             float64   `bun:"downpayment_pct,type:decimal(19,4),nullzero"`
	UseCurrentMonthNetDay      string    `bun:"use_current_month_net_day,type:char,default:('N')"`
	UseCurrentMonthTermsDay    string    `bun:"use_current_month_terms_day,type:char,default:('N')"`
	BillingCycleCutoffDay      int16     `bun:"billing_cycle_cutoff_day,type:smallint,default:(0)"`
	TermsType                  int32     `bun:"terms_type,type:int,default:(1622)"`
	CreditcardRequiredFlag     string    `bun:"creditcard_required_flag,type:char,default:('N')"`
	DiscountInvMastUid         int32     `bun:"discount_inv_mast_uid,type:int,nullzero"`
	CashDiscountEligibleFlag   string    `bun:"cash_discount_eligible_flag,type:char,default:('N')"`
	MerchCreditOnlyFlag        string    `bun:"merch_credit_only_flag,type:char,nullzero"`
	ExcludeDirectShipmentTerms string    `bun:"exclude_direct_shipment_terms,type:char,nullzero"`
	DirectShipmentTermsDesc    string    `bun:"direct_shipment_terms_desc,type:varchar(20),nullzero"`
}
