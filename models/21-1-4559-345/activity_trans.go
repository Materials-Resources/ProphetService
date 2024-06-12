package model

type ActivityTrans struct {
	bun.BaseModel         `bun:"table:activity_trans"`
	ActivityTransNo       string    `bun:"activity_trans_no,type:varchar(10),pk"`
	ActivityId            string    `bun:"activity_id,type:varchar(10)"`
	ContactId             string    `bun:"contact_id,type:varchar(16),nullzero"`
	EntryDate             time.Time `bun:"entry_date,type:datetime"`
	AssignedById          string    `bun:"assigned_by_id,type:varchar(30)"`
	AssignedToId          string    `bun:"assigned_to_id,type:varchar(30),nullzero"`
	CompletedDate         time.Time `bun:"completed_date,type:datetime,nullzero"`
	CompletedFlag         string    `bun:"completed_flag,type:char"`
	CompletedById         string    `bun:"completed_by_id,type:varchar(30),nullzero"`
	Comments              string    `bun:"comments,type:text(2147483647),nullzero"`
	DateCreated           time.Time `bun:"date_created,type:datetime"`
	DateLastModified      time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy      string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	Reminder              string    `bun:"reminder,type:varchar,default:('N')"`
	ReminderTimeOffset    int32     `bun:"reminder_time_offset,type:int,nullzero"`
	ReminderTimeOffsetCd  int32     `bun:"reminder_time_offset_cd,type:int,nullzero"`
	PrivateTask           string    `bun:"private_task,type:varchar,default:('N')"`
	TargetCompleteDate    time.Time `bun:"target_complete_date,type:datetime"`
	Followup              string    `bun:"followup,type:varchar,default:('N')"`
	FollowupCommentCd     int32     `bun:"followup_comment_cd,type:int,default:(300)"`
	TransactionTypeCd     int32     `bun:"transaction_type_cd,type:int,default:(300)"`
	TransactionNo         string    `bun:"transaction_no,type:varchar(255),nullzero"`
	Subject               string    `bun:"subject,type:varchar(255),nullzero"`
	CompanyId             string    `bun:"company_id,type:varchar(8),nullzero"`
	LinkId                float64   `bun:"link_id,type:decimal(19,0),nullzero"`
	LinkTypeCd            int32     `bun:"link_type_cd,type:int,nullzero"`
	HardTouchFlag         string    `bun:"hard_touch_flag,type:char,default:('N')"`
	CreateOutlookTaskFlag string    `bun:"create_outlook_task_flag,type:varchar,default:('N')"`
	DateLastSynched       time.Time `bun:"date_last_synched,type:datetime,nullzero"`
	SerialNumberUid       int32     `bun:"serial_number_uid,type:int,nullzero"`
	AlarmCodeUid          int32     `bun:"alarm_code_uid,type:int,nullzero"`
	ProblemCodeUid        int32     `bun:"problem_code_uid,type:int,nullzero"`
	OutlookTaskUid        string    `bun:"outlook_task_uid,type:varchar(255),nullzero"`
	SyncTaskTypeCd        int32     `bun:"sync_task_type_cd,type:int,default:((2740))"`
	StartDate             time.Time `bun:"start_date,type:datetime,nullzero"`
}
