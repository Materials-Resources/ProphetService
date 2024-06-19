package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type ScheduledJobNotifications struct {
	bun.BaseModel    `bun:"table:scheduled_job_notifications"`
	NotificationUid  int32     `bun:"notification_uid,type:int,autoincrement,pk"`                                                                                                                                                             // UID Coulmn
	NotificationType string    `bun:"notification_type,type:Error: 50000, Severity: -1, State: 1. (Params:). The error is printed in terse mode because there was error during formatting. Tracing, ETW, notifications etc are skipped.\r\n"` // Type of Notification EMAIL,Alert etc.
	NotificationDesc string    `bun:"notification_desc,type:Error: 50000, Severity: -1, State: 1. (Params:). The error is printed in terse mode because there was error during formatting. Tracing, ETW, notifications etc are skipped.\r\n"` // Notification Description
	NotificationName string    `bun:"notification_name,type:Error: 50000, Severity: -1, State: 1. (Params:). The error is printed in terse mode because there was error during formatting. Tracing, ETW, notifications etc are skipped.\r\n"` // Notification Name
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`                                                                                                                                                         // Date and time the record was originally created
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`                                                                                                                                                   // User who created the record
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`                                                                                                                                                   // Date and time the record was modified
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`                                                                                                                                           // User who last changed the record
	DeleteFlag       string    `bun:"delete_flag,type:char(1)"`                                                                                                                                                                               // Delete flag
}
