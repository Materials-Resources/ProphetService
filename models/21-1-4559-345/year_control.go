package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type YearControl struct {
	bun.BaseModel    `bun:"table:year_control"`
	YearControlUid   int32     `bun:"year_control_uid,type:int,pk"`                                  // Internally generated unique identifier.
	CompanyNo        string    `bun:"company_no,type:varchar(8),unique"`                             // Unique code that identifies a company.
	Year             float64   `bun:"year,type:decimal(4,0),unique"`                                 // Indicates the year.
	ClosedFlag       string    `bun:"closed_flag,type:char(1),default:('N')"`                        // Indicates that this record is closed.
	DeleteFlag       string    `bun:"delete_flag,type:char(1),default:('N')"`                        // Indicates whether this record is logically deleted
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`                // Indicates the date/time this record was created.
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`          // Indicates the date/time this record was last modified.
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"` // ID of the user who last maintained this record
	GlRollupFlag     string    `bun:"gl_rollup_flag,type:char(1),default:('N')"`                     // Indicates whether this periods GL records have been rolled up (and deleted).
}
