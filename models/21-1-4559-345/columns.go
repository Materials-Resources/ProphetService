package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type Columns struct {
	bun.BaseModel     `bun:"table:columns"`
	FinReportId       string    `bun:"fin_report_id,type:varchar(15),pk"`                         // Financial report associated with this columns record
	ColumnNo          float64   `bun:"column_no,type:decimal(19,0),pk"`                           // Column number in the financial report spreadsheet
	BaseYearOffset    float64   `bun:"base_year_offset,type:decimal(4,0)"`                        // Number of years "away from" the base year entered in Update Financial Worksheets
	BasePeriodOffset  float64   `bun:"base_period_offset,type:decimal(3,0)"`                      // Number of years "away from" the base period entered in Update Financial Worksheets
	DateCreated       time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified  time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy  string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	ColumnsUid        int32     `bun:"columns_uid,type:int,autoincrement,unique"`                 // Unique identity column for columns
	BaseQuarterOffset float64   `bun:"base_quarter_offset,type:decimal(4,0),default:((0))"`       // Quarter offset used for quarterly financial statements.
}
