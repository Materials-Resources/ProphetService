package model

type BusinessRuleEventExtdInfo struct {
	bun.BaseModel                `bun:"table:business_rule_event_extd_info"`
	BusinessRuleEventExtdInfoUid int32 `bun:"business_rule_event_extd_info_uid,type:int,pk,identity"`
	BusinessRuleEventUid         int32 `bun:"business_rule_event_uid,type:int"`
	TriggeredFrom                `bun:"triggered_from,type:nvarchar(4000)"`
	Keys                         `bun:"keys,type:nvarchar(255)"`
	Context                      `bun:"context,type:nvarchar(4000)"`
	SpecialReturnValues          `bun:"special_return_values,type:nvarchar(255)"`
	SeeAlso                      `bun:"see_also,type:nvarchar(255)"`
	DateCreated                  time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                    string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified             time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy             string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
