package model

import (
	"github.com/uptrace/bun"
	"time"
)

type ActivityReminder struct {
	bun.BaseModel       `bun:"table:activity_reminder"`
	ActivityReminderUid int32     `bun:"activity_reminder_uid,type:int,pk"`
	ActivityTransNo     string    `bun:"activity_trans_no,type:varchar(10)"`
	AssignedToId        string    `bun:"assigned_to_id,type:varchar(30)"`
	ReminderDate        time.Time `bun:"reminder_date,type:datetime"`
	Displaying          string    `bun:"displaying,type:varchar,default:('N')"`
	ReminderStatusCd    int32     `bun:"reminder_status_cd,type:int,default:(1442)"`
	RowStatusFlag       int32     `bun:"row_status_flag,type:int"`
	DateCreated         time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy           string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified    time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy    string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
