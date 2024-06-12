package model

type CommissionScheduleDetail struct {
	bun.BaseModel               `bun:"table:commission_schedule_detail"`
	CompanyId                   string    `bun:"company_id,type:varchar(8),pk"`
	CommissionScheduleId        string    `bun:"commission_schedule_id,type:varchar(8),pk"`
	CommissionRuleId            string    `bun:"commission_rule_id,type:varchar(8),pk"`
	SortOrder                   float64   `bun:"sort_order,type:decimal(19,0),pk"`
	Inactive                    string    `bun:"inactive,type:char"`
	DeleteFlag                  string    `bun:"delete_flag,type:char"`
	DateCreated                 time.Time `bun:"date_created,type:datetime"`
	DateLastModified            time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy            string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	CommissionScheduleDetailUid int32     `bun:"commission_schedule_detail_uid,type:int,identity"`
}
