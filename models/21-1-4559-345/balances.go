package model

import (
	"github.com/uptrace/bun"
	"time"
)

type Balances struct {
	bun.BaseModel            `bun:"table:balances"`
	CompanyNo                string    `bun:"company_no,type:varchar(8),pk"`
	AccountNo                string    `bun:"account_no,type:varchar(32),pk"`
	Period                   float64   `bun:"period,type:decimal(3,0),pk"`
	YearForPeriod            float64   `bun:"year_for_period,type:decimal(4,0),pk"`
	CumulativeBalance        float64   `bun:"cumulative_balance,type:decimal(19,4)"`
	PeriodBalance            float64   `bun:"period_balance,type:decimal(19,4)"`
	Budget1                  float64   `bun:"budget_1,type:decimal(19,4)"`
	Budget2                  float64   `bun:"budget_2,type:decimal(19,4)"`
	Budget3                  float64   `bun:"budget_3,type:decimal(19,4)"`
	DeleteFlag               string    `bun:"delete_flag,type:char"`
	DateCreated              time.Time `bun:"date_created,type:datetime"`
	DateLastModified         time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy         string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	EncumberedBalance        float64   `bun:"encumbered_balance,type:decimal(19,4),default:(0)"`
	EncumberedThisPeriod     float64   `bun:"encumbered_this_period,type:decimal(19,4),default:(0)"`
	DateBudgetChanged        time.Time `bun:"date_budget_changed,type:datetime,default:('01/01/1980')"`
	ForeignPeriodBalance     float64   `bun:"foreign_period_balance,type:decimal(19,4),nullzero"`
	ForeignCumulativeBalance float64   `bun:"foreign_cumulative_balance,type:decimal(19,4),nullzero"`
	CurrencyId               float64   `bun:"currency_id,type:decimal(19,0),pk"`
	CumulativeBudget1        float64   `bun:"cumulative_budget_1,type:decimal(19,2),default:((0))"`
	CumulativeBudget2        float64   `bun:"cumulative_budget_2,type:decimal(19,2),default:((0))"`
	CumulativeBudget3        float64   `bun:"cumulative_budget_3,type:decimal(19,2),default:((0))"`
}
