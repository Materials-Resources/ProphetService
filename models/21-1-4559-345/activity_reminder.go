package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type ActivityReminder struct {
	bun.BaseModel       `bun:"table:activity_reminder"`
	ActivityReminderUid int32     `bun:"activity_reminder_uid,type:int,pk"`                            // Unique record identifier.
	ActivityTransNo     string    `bun:"activity_trans_no,type:varchar(10)"`                           // Identifier for activity trans
	AssignedToId        string    `bun:"assigned_to_id,type:varchar(30)"`                              // User id assigned to complete the activity
	ReminderDate        time.Time `bun:"reminder_date,type:datetime"`                                  // Reminder time
	Displaying          string    `bun:"displaying,type:varchar(1),default:('N')"`                     // Indicates message is currently displaying
	ReminderStatusCd    int32     `bun:"reminder_status_cd,type:int,default:(1442)"`                   // Gives historical detail about the reminder date
	RowStatusFlag       int32     `bun:"row_status_flag,type:int"`                                     // Indicates current record status
	DateCreated         time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy           string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified    time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy    string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
