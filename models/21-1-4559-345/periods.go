package model

type Periods struct {
	bun.BaseModel        `bun:"table:periods"`
	CompanyNo            string    `bun:"company_no,type:varchar(8),pk"`
	Period               float64   `bun:"period,type:decimal(3,0),pk"`
	YearForPeriod        float64   `bun:"year_for_period,type:decimal(4,0),pk"`
	PeriodClosed         string    `bun:"period_closed,type:char"`
	DateCreated          time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	DateLastModified     time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy     string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	BeginningDate        time.Time `bun:"beginning_date,type:datetime,nullzero"`
	EndingDate           time.Time `bun:"ending_date,type:datetime,nullzero"`
	AdjustmentPeriodFlag string    `bun:"adjustment_period_flag,type:char,default:('N')"`
	GlRollupFlag         string    `bun:"gl_rollup_flag,type:char,default:('N')"`
	PeriodReopened       string    `bun:"period_reopened,type:char,nullzero"`
}
