package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type ActivityTrans struct {
	bun.BaseModel         `bun:"table:activity_trans"`
	ActivityTransNo       string     `bun:"activity_trans_no,type:varchar(10),pk"`                     // Automatically system generated id
	ActivityId            string     `bun:"activity_id,type:varchar(10)"`                              // This is the system generated unique identifier for an activity trans.
	ContactId             *string    `bun:"contact_id,type:varchar(16)"`                               // What contact deals with this sales representative?
	EntryDate             time.Time  `bun:"entry_date,type:datetime"`                                  // Date indicating when this note was entered.
	AssignedById          string     `bun:"assigned_by_id,type:varchar(30)"`                           // User id that assigned the activity
	AssignedToId          *string    `bun:"assigned_to_id,type:varchar(30)"`                           // user id assigned to complete the activity
	CompletedDate         *time.Time `bun:"completed_date,type:datetime"`                              // Date the activity was completed -  or if the complet
	CompletedFlag         string     `bun:"completed_flag,type:char(1)"`                               // Check box showing whether activity has been comple
	CompletedById         *string    `bun:"completed_by_id,type:varchar(30)"`                          // User id who completed the activity
	Comments              *string    `bun:"comments,type:text"`                                        // User defined comments about the activity
	DateCreated           time.Time  `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified      time.Time  `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy      string     `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	Reminder              string     `bun:"reminder,type:varchar(1),default:('N')"`                    // Indicates a reminder is desired for this task.
	ReminderTimeOffset    *int32     `bun:"reminder_time_offset,type:int"`                             // Time before target complete date reminder should display
	ReminderTimeOffsetCd  *int32     `bun:"reminder_time_offset_cd,type:int"`                          // Qualifies type of time the offset refers to
	PrivateTask           string     `bun:"private_task,type:varchar(1),default:('N')"`                // Indicates task is reserved for assigned users
	TargetCompleteDate    time.Time  `bun:"target_complete_date,type:datetime"`                        // Date by which we would like this task completed
	Followup              string     `bun:"followup,type:varchar(1),default:('N')"`                    // Followup requested when task completes
	FollowupCommentCd     int32      `bun:"followup_comment_cd,type:int,default:(300)"`                // Describes if comments will transfer to followup task
	TransactionTypeCd     int32      `bun:"transaction_type_cd,type:int,default:(300)"`                // Indicates type of transaction referenced
	TransactionNo         *string    `bun:"transaction_no,type:varchar(255)"`                          // Transaction reference
	Subject               *string    `bun:"subject,type:varchar(255)"`                                 // Overview of task
	CompanyId             *string    `bun:"company_id,type:varchar(8)"`                                // Unique code that identifies a company.
	LinkId                *float64   `bun:"link_id,type:decimal(19,0)"`                                // Unique code that elaborates the role the contact_id is playing for this task
	LinkTypeCd            *int32     `bun:"link_type_cd,type:int"`                                     // Unique code which describes the link_id
	HardTouchFlag         string     `bun:"hard_touch_flag,type:char(1),default:('N')"`                // Indicates that the task involves direct contact with the customer, supplier, etc
	CreateOutlookTaskFlag *string    `bun:"create_outlook_task_flag,type:varchar(1),default:('N')"`    // Indicates task is to be created in outlook
	DateLastSynched       *time.Time `bun:"date_last_synched,type:datetime"`                           // Indicates the date/time this record was last synched with outlook
	SerialNumberUid       *int32     `bun:"serial_number_uid,type:int"`                                // Indicates the serial number specific for this activity - FK to serial_number.
	AlarmCodeUid          *int32     `bun:"alarm_code_uid,type:int"`                                   // Alarm code for task.
	ProblemCodeUid        *int32     `bun:"problem_code_uid,type:int"`                                 // Problem code for task.
	OutlookTaskUid        *string    `bun:"outlook_task_uid,type:varchar(255)"`                        // Unique identifier for an outlook task.
	SyncTaskTypeCd        *int32     `bun:"sync_task_type_cd,type:int,default:((2740))"`               // Code value to determine whether to create a task or appointment.
	StartDate             *time.Time `bun:"start_date,type:datetime"`                                  // Start Date for the task
}
