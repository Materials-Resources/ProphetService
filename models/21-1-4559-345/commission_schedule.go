package model

type CommissionSchedule struct {
	bun.BaseModel          `bun:"table:commission_schedule"`
	CompanyId              string    `bun:"company_id,type:varchar(8),pk"`
	CommissionScheduleId   string    `bun:"commission_schedule_id,type:varchar(8),pk"`
	CommissionScheduleDesc string    `bun:"commission_schedule_desc,type:varchar(40)"`
	CommissionScheduleType string    `bun:"commission_schedule_type,type:char"`
	Inactive               string    `bun:"inactive,type:char"`
	DeleteFlag             string    `bun:"delete_flag,type:char"`
	DateCreated            time.Time `bun:"date_created,type:datetime"`
	DateLastModified       time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy       string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	AppliesTo              string    `bun:"applies_to,type:char"`
	CommissionScheduleUid  int32     `bun:"commission_schedule_uid,type:int,identity"`
}
