package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type AlertImplementation struct {
	bun.BaseModel           `bun:"table:alert_implementation"`
	AlertImplementationUid  int32     `bun:"alert_implementation_uid,type:int,pk"`               // Unique Identifier for record.
	AlertTypeUid            int32     `bun:"alert_type_uid,type:int,unique"`                     // Unique Identifier from alert_type table.
	AlertImplementationName string    `bun:"alert_implementation_name,type:varchar(255),unique"` // Name of alert implmentation query
	ActivationDate          time.Time `bun:"activation_date,type:datetime"`                      // The date this note will be activated.
	ExpirationDate          time.Time `bun:"expiration_date,type:datetime"`                      // The date this note will expire.
	DateCreated             time.Time `bun:"date_created,type:datetime"`                         // Indicates the date/time this record was created.
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime"`                   // Indicates the date/time this record was last modified.
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(30)"`                // ID of the user who last maintained this record
	RowStatusFlag           int32     `bun:"row_status_flag,type:int"`                           // Indicates current record status.
	WhereClause             string    `bun:"where_clause,type:text,nullzero"`                    // The where clause that will be used for a specific alert.
	ExecutionModeCd         int32     `bun:"execution_mode_cd,type:int"`                         // Code that indicates how this alert will execute - Scheduled or Immediate.
	NextExecutionStartPoint time.Time `bun:"next_execution_start_point,type:datetime,nullzero"`  // Indicates the next time this alert will start.
	JobId                   string    `bun:"job_id,type:varchar(255)"`                           // Indentifier for the job.  Currently unused. Now this is stored in the alert_type table.
	LastExecutionError      int32     `bun:"last_execution_error,type:int,nullzero"`             // Executing a string does not return a value.
	EmailNotificationFlag   string    `bun:"email_notification_flag,type:char(1),default:('Y')"` // Indicates whether an email will be generated when the alert is triggered.
	TaskNotificationFlag    string    `bun:"task_notification_flag,type:char(1),default:('N')"`  // Indicates whether a task will be generated when the alert is triggered.
}
