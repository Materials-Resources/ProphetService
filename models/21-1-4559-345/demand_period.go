package model

import (
	"github.com/uptrace/bun"
	"time"
)

type DemandPeriod struct {
	bun.BaseModel      `bun:"table:demand_period"`
	DemandPeriodUid    int32     `bun:"demand_period_uid,type:int,pk"`
	CompanyId          string    `bun:"company_id,type:varchar(8)"`
	Period             int16     `bun:"period,type:smallint"`
	YearForPeriod      int16     `bun:"year_for_period,type:smallint"`
	BeginningDate      time.Time `bun:"beginning_date,type:datetime"`
	EndingDate         time.Time `bun:"ending_date,type:datetime"`
	DateCreated        time.Time `bun:"date_created,type:datetime"`
	DateLastModified   time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy   string    `bun:"last_maintained_by,type:varchar(30)"`
	RowStatusFlag      int16     `bun:"row_status_flag,type:smallint"`
	ComputedYearPeriod int32     `bun:"computed_year_period,type:int,nullzero"`
}
