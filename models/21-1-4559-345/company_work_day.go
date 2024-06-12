package model

import (
	"github.com/uptrace/bun"
	"time"
)

type CompanyWorkDay struct {
	bun.BaseModel     `bun:"table:company_work_day"`
	CompanyWorkDayUid int32     `bun:"company_work_day_uid,type:int,pk"`
	CompanyId         string    `bun:"company_id,type:varchar(8),nullzero"`
	WorkDay           int32     `bun:"work_day,type:int"`
	ActiveFlag        string    `bun:"active_flag,type:char,default:('N')"`
	StartTime         time.Time `bun:"start_time,type:datetime"`
	EndTime           time.Time `bun:"end_time,type:datetime"`
	DateCreated       time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy         string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified  time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy  string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
