package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type DemandPeriod struct {
	bun.BaseModel      `bun:"table:demand_period"`
	DemandPeriodUid    int32     `bun:"demand_period_uid,type:int,pk"`       // unique identifier for demand period records
	CompanyId          string    `bun:"company_id,type:varchar(8)"`          // Unique code that identifies a company.
	Period             int16     `bun:"period,type:smallint"`                // Which period does this quota apply to?
	YearForPeriod      int16     `bun:"year_for_period,type:smallint"`       // What year does the period belong to?
	BeginningDate      time.Time `bun:"beginning_date,type:datetime"`        // What is the beginning date for this period?
	EndingDate         time.Time `bun:"ending_date,type:datetime"`           // When is the end of this period?
	DateCreated        time.Time `bun:"date_created,type:datetime"`          // Indicates the date/time this record was created.
	DateLastModified   time.Time `bun:"date_last_modified,type:datetime"`    // Indicates the date/time this record was last modified.
	LastMaintainedBy   string    `bun:"last_maintained_by,type:varchar(30)"` // ID of the user who last maintained this record
	RowStatusFlag      int16     `bun:"row_status_flag,type:smallint"`       // Indicates current record status.
	ComputedYearPeriod *int32    `bun:"computed_year_period,type:int"`       // computed field made up of year_for_period + year
}
