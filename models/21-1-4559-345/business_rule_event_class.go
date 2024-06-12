package model

type BusinessRuleEventClass struct {
	bun.BaseModel             `bun:"table:business_rule_event_class"`
	BusinessRuleEventClassUid int32     `bun:"business_rule_event_class_uid,type:int,pk,identity"`
	BusinessRuleEventUid      int32     `bun:"business_rule_event_uid,type:int"`
	RuleClassName             string    `bun:"rule_class_name,type:varchar(255)"`
	RunTypeCd                 int32     `bun:"run_type_cd,type:int"`
	ConfigurationId           int32     `bun:"configuration_id,type:int,nullzero"`
	SequenceNo                int32     `bun:"sequence_no,type:int"`
	DateCreated               time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                 string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified          time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy          string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
