package model

import (
	"github.com/uptrace/bun"
	"time"
)

type ScheduledJobHistory struct {
	bun.BaseModel          `bun:"table:scheduled_job_history"`
	ScheduledJobHistoryUid int32     `bun:"scheduled_job_history_uid,type:int,pk,identity"`
	ScheduledJobUid        int32     `bun:"scheduled_job_uid,type:int"`
	JobRunAtDate           time.Time `bun:"job_run_at_date,type:datetime"`
	JobRunStatus           string    `bun:"job_run_status,type:varchar(50)"`
	Message                string    `bun:"message,type:nvarchar"`
	DateCreated            time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy              string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified       time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy       string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
