package model

type BusinessRuleEvent struct {
	bun.BaseModel        `bun:"table:business_rule_event"`
	BusinessRuleEventUid int32     `bun:"business_rule_event_uid,type:int,pk,identity"`
	PublishedEventName   string    `bun:"published_event_name,type:varchar(255)"`
	DisplayName          string    `bun:"display_name,type:varchar(255)"`
	Description          string    `bun:"description,type:varchar(255)"`
	InternalOnlyFlag     string    `bun:"internal_only_flag,type:char,default:('N')"`
	DateCreated          time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy            string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified     time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy     string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	AllowNewRowsFlag     string    `bun:"allow_new_rows_flag,type:char,default:('N')"`
}
