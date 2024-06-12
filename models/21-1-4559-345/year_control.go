package model

import (
	"github.com/uptrace/bun"
	"time"
)

type YearControl struct {
	bun.BaseModel    `bun:"table:year_control"`
	YearControlUid   int32     `bun:"year_control_uid,type:int,pk"`
	CompanyNo        string    `bun:"company_no,type:varchar(8)"`
	Year             float64   `bun:"year,type:decimal(4,0)"`
	ClosedFlag       string    `bun:"closed_flag,type:char,default:('N')"`
	DeleteFlag       string    `bun:"delete_flag,type:char,default:('N')"`
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"`
	GlRollupFlag     string    `bun:"gl_rollup_flag,type:char,default:('N')"`
}
