package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type Balances struct {
	bun.BaseModel            `bun:"table:balances"`
	CompanyNo                string    `bun:"company_no,type:varchar(8),pk"`                             // Unique code that identifies a company.
	AccountNo                string    `bun:"account_no,type:varchar(32),pk"`                            // The account which the balances record applies to.
	Period                   float64   `bun:"period,type:decimal(3,0),pk"`                               // The period which the balance is for.
	YearForPeriod            float64   `bun:"year_for_period,type:decimal(4,0),pk"`                      // The year which the balance is for.
	CumulativeBalance        float64   `bun:"cumulative_balance,type:decimal(19,4)"`                     // Sum of the period balances for current and previous periods for a particular account.  In the case of revenue or expense accts, this column is the sum of the period balances for the current  year only.
	PeriodBalance            float64   `bun:"period_balance,type:decimal(19,4)"`                         // The balance of any one account for the particular year/period.
	Budget1                  float64   `bun:"budget_1,type:decimal(19,4)"`                               // The budgeted amount for an account for a given period/year.
	Budget2                  float64   `bun:"budget_2,type:decimal(19,4)"`                               // The budgeted amount for an account for a given period/year.
	Budget3                  float64   `bun:"budget_3,type:decimal(19,4)"`                               // The budgeted amount for an account for a given period/year.
	DeleteFlag               string    `bun:"delete_flag,type:char(1)"`                                  // Indicates whether this record is logically deleted
	DateCreated              time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified         time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy         string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	EncumberedBalance        float64   `bun:"encumbered_balance,type:decimal(19,4),default:(0)"`         // Sum of transactions posted to encumbered account for a particular year-to-date.
	EncumberedThisPeriod     float64   `bun:"encumbered_this_period,type:decimal(19,4),default:(0)"`     // Sum of transactions posted to encumbered account for a single period.
	DateBudgetChanged        time.Time `bun:"date_budget_changed,type:datetime,default:('01/01/1980')"`  // Indicates the date when the budget last changed.
	ForeignPeriodBalance     float64   `bun:"foreign_period_balance,type:decimal(19,4),nullzero"`        // Stores foreign amount for period per currency
	ForeignCumulativeBalance float64   `bun:"foreign_cumulative_balance,type:decimal(19,4),nullzero"`    // Stores accumulated foreign amount per currency
	CurrencyId               float64   `bun:"currency_id,type:decimal(19,0),pk"`                         // Indicates currency of foreign amounts
	CumulativeBudget1        float64   `bun:"cumulative_budget_1,type:decimal(19,2),default:((0))"`      // Cumulative balance of budget 1.
	CumulativeBudget2        float64   `bun:"cumulative_budget_2,type:decimal(19,2),default:((0))"`      // Cumulative balance of budget 2.
	CumulativeBudget3        float64   `bun:"cumulative_budget_3,type:decimal(19,2),default:((0))"`      // Cumulative balance of budget 3.
}
