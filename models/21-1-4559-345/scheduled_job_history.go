package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type ScheduledJobHistory struct {
	bun.BaseModel          `bun:"table:scheduled_job_history"`
	ScheduledJobHistoryUid int32     `bun:"scheduled_job_history_uid,type:int,autoincrement,scanonly,pk"`                                                                                                                                 // identity & auto-incremented column
	ScheduledJobUid        int32     `bun:"scheduled_job_uid,type:int"`                                                                                                                                                                   // foreign key to scheduled_job
	JobRunAtDate           time.Time `bun:"job_run_at_date,type:datetime"`                                                                                                                                                                // last time when job finished it's task
	JobRunStatus           string    `bun:"job_run_status,type:varchar(50)"`                                                                                                                                                              // job status, failed, cancelled or success
	Message                string    `bun:"message,type:Error: 50000, Severity: -1, State: 1. (Params:). The error is printed in terse mode because there was error during formatting. Tracing, ETW, notifications etc are skipped.\r\n"` // message related to job running status. In case of exception it will have complete stach trace information including exception details. Hence this is required to be nvarchar(max)
	DateCreated            time.Time `bun:"date_created,type:datetime,default:(getdate())"`                                                                                                                                               // Date and time the record was originally created
	CreatedBy              string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`                                                                                                                                         // User who created the record
	DateLastModified       time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`                                                                                                                                         // Date and time the record was modified
	LastMaintainedBy       string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`                                                                                                                                 // User who last changed the record
}
