package model

import (
	"github.com/uptrace/bun"
	"time"
)

type AlertImplementation struct {
	bun.BaseModel           `bun:"table:alert_implementation"`
	AlertImplementationUid  int32     `bun:"alert_implementation_uid,type:int,pk"`
	AlertTypeUid            int32     `bun:"alert_type_uid,type:int"`
	AlertImplementationName string    `bun:"alert_implementation_name,type:varchar(255)"`
	ActivationDate          time.Time `bun:"activation_date,type:datetime"`
	ExpirationDate          time.Time `bun:"expiration_date,type:datetime"`
	DateCreated             time.Time `bun:"date_created,type:datetime"`
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(30)"`
	RowStatusFlag           int32     `bun:"row_status_flag,type:int"`
	WhereClause             string    `bun:"where_clause,type:text(2147483647),nullzero"`
	ExecutionModeCd         int32     `bun:"execution_mode_cd,type:int"`
	NextExecutionStartPoint time.Time `bun:"next_execution_start_point,type:datetime,nullzero"`
	JobId                   string    `bun:"job_id,type:varchar(255)"`
	LastExecutionError      int32     `bun:"last_execution_error,type:int,nullzero"`
	EmailNotificationFlag   string    `bun:"email_notification_flag,type:char,default:('Y')"`
	TaskNotificationFlag    string    `bun:"task_notification_flag,type:char,default:('N')"`
}
