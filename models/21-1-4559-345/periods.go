package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type Periods struct {
	bun.BaseModel        `bun:"table:periods"`
	CompanyNo            string    `bun:"company_no,type:varchar(8),pk"`                             // Unique code that identifies a company.
	Period               float64   `bun:"period,type:decimal(3,0),pk"`                               // In which period did the inventory transfer occur?
	YearForPeriod        float64   `bun:"year_for_period,type:decimal(4,0),pk"`                      // What year does the period belong to?
	PeriodClosed         string    `bun:"period_closed,type:char(1)"`                                // Indicate whether to close or open a period
	DateCreated          time.Time `bun:"date_created,type:datetime,default:(getdate())"`            // Indicates the date/time this record was created.
	DateLastModified     time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`      // Indicates the date/time this record was last modified.
	LastMaintainedBy     string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	BeginningDate        time.Time `bun:"beginning_date,type:datetime,nullzero"`                     // What is the beginning date for this period?
	EndingDate           time.Time `bun:"ending_date,type:datetime,nullzero"`                        // When is the end of this period?
	AdjustmentPeriodFlag string    `bun:"adjustment_period_flag,type:char(1),default:('N')"`         // Indicates whether this period is an adjustment period.
	GlRollupFlag         string    `bun:"gl_rollup_flag,type:char(1),default:('N')"`                 // Indicates whether this periods GL records have been rolled up (and deleted).
	PeriodReopened       string    `bun:"period_reopened,type:char(1),nullzero"`                     // A flag to determine whether a period has been reopened
}
