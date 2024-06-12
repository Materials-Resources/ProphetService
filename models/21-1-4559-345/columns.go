package model

import (
	"github.com/uptrace/bun"
	"time"
)

type Columns struct {
	bun.BaseModel     `bun:"table:columns"`
	FinReportId       string    `bun:"fin_report_id,type:varchar(15),pk"`
	ColumnNo          float64   `bun:"column_no,type:decimal(19,0),pk"`
	BaseYearOffset    float64   `bun:"base_year_offset,type:decimal(4,0)"`
	BasePeriodOffset  float64   `bun:"base_period_offset,type:decimal(3,0)"`
	DateCreated       time.Time `bun:"date_created,type:datetime"`
	DateLastModified  time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy  string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	ColumnsUid        int32     `bun:"columns_uid,type:int,identity"`
	BaseQuarterOffset float64   `bun:"base_quarter_offset,type:decimal(4,0),default:((0))"`
}
