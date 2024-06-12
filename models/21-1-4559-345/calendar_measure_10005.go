package model

import (
	"github.com/uptrace/bun"
	"time"
)

type CalendarMeasure10005 struct {
	bun.BaseModel       `bun:"table:calendar_measure_10005"`
	CalendarMeasureUid  int32     `bun:"calendar_measure_uid,type:int,pk"`
	CalendarMeasureName string    `bun:"calendar_measure_name,type:varchar(10)"`
	DateCreated         time.Time `bun:"date_created,type:datetime"`
	DateLastModified    time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy    string    `bun:"last_maintained_by,type:varchar(30)"`
	CalendarMeasureCode string    `bun:"calendar_measure_code,type:char(3)"`
}
