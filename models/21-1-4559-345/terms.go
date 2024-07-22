package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type Terms struct {
	bun.BaseModel              `bun:"table:terms"`
	TermsId                    string    `bun:"terms_id,type:varchar(2),pk"`                                   // Terms identification number.
	TermsDesc                  string    `bun:"terms_desc,type:varchar(20)"`                                   // Description of terms.
	DiscountPct                *float64  `bun:"discount_pct,type:decimal(19,4)"`                               // Percentage of terms discount.
	DiscountDays               *float64  `bun:"discount_days,type:decimal(3,0)"`                               // Number of days after invoice date which terms discount is available.
	NetDays                    *float64  `bun:"net_days,type:decimal(3,0)"`                                    // Number of days after invoice date is the invoices net due date.
	DeleteFlag                 string    `bun:"delete_flag,type:char(1)"`                                      // Indicates whether this record is logically deleted
	DateCreated                time.Time `bun:"date_created,type:datetime"`                                    // Indicates the date/time this record was created.
	DateLastModified           time.Time `bun:"date_last_modified,type:datetime"`                              // Indicates the date/time this record was last modified.
	LastMaintainedBy           string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"` // ID of the user who last maintained this record
	DayOfMonth                 *float64  `bun:"day_of_month,type:decimal(2,0)"`                                // Indicates day of month.
	Months                     *float64  `bun:"months,type:decimal(2,0)"`                                      // Indicates number of months.
	TermsDayOfMonth            *float64  `bun:"terms_day_of_month,type:decimal(2,0)"`                          // Indicates which day of the month is the terms due date.
	TermsMonths                *float64  `bun:"terms_months,type:decimal(2,0)"`                                // Indicates additional months for calculating terms due date.
	NetDayOfMonth2             *float64  `bun:"net_day_of_month2,type:decimal(2,0)"`                           // Indicates day of second half of month .
	TermsDayOfMonth2           *float64  `bun:"terms_day_of_month2,type:decimal(2,0)"`                         // Indicates which day in the second half of the month is the terms due date.
	DownpaymentPct             *float64  `bun:"downpayment_pct,type:decimal(19,4)"`                            // Indicates the downpayment percentage.
	UseCurrentMonthNetDay      string    `bun:"use_current_month_net_day,type:char(1),default:('N')"`
	UseCurrentMonthTermsDay    string    `bun:"use_current_month_terms_day,type:char(1),default:('N')"`
	BillingCycleCutoffDay      int16     `bun:"billing_cycle_cutoff_day,type:smallint,default:(0)"`     // Indicates the billing cycle cutoff day for the Day of Month payment terms.
	TermsType                  int32     `bun:"terms_type,type:int,default:(1622)"`                     // Indicates use preset dates or user defined dates to split the month
	CreditcardRequiredFlag     string    `bun:"creditcard_required_flag,type:char(1),default:('N')"`    // Indicates if a creditcard required for these terms.
	DiscountInvMastUid         *int32    `bun:"discount_inv_mast_uid,type:int"`                         // Identifies the item to be used for cash discounts
	CashDiscountEligibleFlag   string    `bun:"cash_discount_eligible_flag,type:char(1),default:('N')"` // Indicates if term type is eligible for cash discount
	MerchCreditOnlyFlag        *string   `bun:"merch_credit_only_flag,type:char(1)"`                    // Customers with this term can only use payment type merchandise credit in RMA
	ExcludeDirectShipmentTerms *string   `bun:"exclude_direct_shipment_terms,type:char(1)"`             // Flag to indicate that terms will be excluded from Direct Shipment invoice calculations.
	DirectShipmentTermsDesc    *string   `bun:"direct_shipment_terms_desc,type:varchar(20)"`            // Terms description to use when Direct Shipments are configured to exclude terms.
}
