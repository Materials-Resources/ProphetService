package model

type BusinessRuleDataElement struct {
	bun.BaseModel              `bun:"table:business_rule_data_element"`
	BusinessRuleDataElementUid int32     `bun:"business_rule_data_element_uid,type:int,pk,identity"`
	BusinessRuleUid            int32     `bun:"business_rule_uid,type:int"`
	FieldName                  string    `bun:"field_name,type:varchar(255)"`
	ClassName                  string    `bun:"class_name,type:varchar(255),nullzero"`
	DateCreated                time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                  string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified           time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy           string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	FieldAlias                 string    `bun:"field_alias,type:varchar(255),nullzero"`
}
