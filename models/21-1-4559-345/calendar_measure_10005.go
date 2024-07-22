package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type CalendarMeasure10005 struct {
	bun.BaseModel       `bun:"table:calendar_measure_10005"`
	CalendarMeasureUid  int32     `bun:"calendar_measure_uid,type:int,pk"`       // Unique internal record identifier - not visible to the user
	CalendarMeasureName string    `bun:"calendar_measure_name,type:varchar(10)"` // Displayed timeframe identifier (Days, Weeks, Months, Years)
	DateCreated         time.Time `bun:"date_created,type:datetime"`             // Indicates the date/time this record was created.
	DateLastModified    time.Time `bun:"date_last_modified,type:datetime"`       // Indicates the date/time this record was last modified.
	LastMaintainedBy    string    `bun:"last_maintained_by,type:varchar(30)"`    // ID of the user who last maintained this record
	CalendarMeasureCode string    `bun:"calendar_measure_code,type:char(3)"`     // Internal code for timeframe identifier
}
