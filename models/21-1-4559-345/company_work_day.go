package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type CompanyWorkDay struct {
	bun.BaseModel     `bun:"table:company_work_day"`
	CompanyWorkDayUid int32     `bun:"company_work_day_uid,type:int,pk"`                             // Unique ID for the Work Day
	CompanyId         string    `bun:"company_id,type:varchar(8),nullzero"`                          // Company ID
	WorkDay           int32     `bun:"work_day,type:int"`                                            // Day of week (1 is Sunday, 7 is Saturday)
	ActiveFlag        string    `bun:"active_flag,type:char(1),default:('N')"`                       // Flag to indicate whether given day is a work day
	StartTime         time.Time `bun:"start_time,type:datetime"`                                     // Start of work day
	EndTime           time.Time `bun:"end_time,type:datetime"`                                       // End of work day
	DateCreated       time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy         string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified  time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy  string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
