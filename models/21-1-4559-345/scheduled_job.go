package model

type ScheduledJob struct {
	bun.BaseModel              `bun:"table:scheduled_job"`
	CompanyId                  string `bun:"company_id,type:varchar(8),nullzero"`
	Name                       string `bun:"name,type:varchar(255)"`
	Description                string `bun:"description,type:varchar(255)"`
	Type                       string `bun:"type,type:varchar(255)"`
	JobConfig                  `bun:"job_config,type:nvarchar,nullzero"`
	StartDate                  time.Time `bun:"start_date,type:datetime"`
	EndDate                    time.Time `bun:"end_date,type:datetime,nullzero"`
	ScheduledIntervalInminutes int32     `bun:"scheduled_interval_inminutes,type:int"`
	ActiveFlag                 string    `bun:"active_flag,type:char,default:('Y')"`
	RunningFlag                string    `bun:"running_flag,type:char,default:('N')"`
	LastRunDate                time.Time `bun:"last_run_date,type:datetime,nullzero"`
	DeleteFlag                 string    `bun:"delete_flag,type:char,default:('N')"`
	DateCreated                time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                  string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified           time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy           string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	ScheduledIntervalInhours   int32     `bun:"scheduled_interval_inhours,type:int,default:((0))"`
	ScheduledIntervalInseconds int32     `bun:"scheduled_interval_inseconds,type:int,default:((0))"`
	RepeatDays                 string    `bun:"repeat_days,type:char(15),nullzero"`
	LastRunStatus              string    `bun:"last_run_status,type:varchar(50),nullzero"`
	ScheduledJobUid            int32     `bun:"scheduled_job_uid,type:int,pk"`
	SchedulerIdentifier        `bun:"scheduler_identifier,type:uniqueidentifier,nullzero"`
	LastPingDate               time.Time `bun:"last_ping_date,type:datetime,nullzero"`
	FirstRunDate               time.Time `bun:"first_run_date,type:datetime,nullzero"`
	RecurrenceType             int32     `bun:"recurrence_type,type:int,default:((0))"`
	RecurrenceIndex            int32     `bun:"recurrence_index,type:int,default:((0))"`
	RecurrenceInterval         int32     `bun:"recurrence_interval,type:int,default:((0))"`
	RecurrenceMonth            int32     `bun:"recurrence_month,type:int,default:((0))"`
	RecurrenceDom              int32     `bun:"recurrence_dom,type:int,default:((0))"`
	ConsecutiveFailureCounter  int32     `bun:"consecutive_failure_counter,type:int,default:((0))"`
	HistoryLevel               int32     `bun:"history_level,type:int,default:((1664))"`
	MaxRetryAttempts           int32     `bun:"max_retry_attempts,type:int,nullzero"`
	RunOnce                    string    `bun:"run_once,type:char,default:('N')"`
}
