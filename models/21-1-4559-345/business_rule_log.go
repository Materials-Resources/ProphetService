package model

type BusinessRuleLog struct {
	bun.BaseModel      `bun:"table:business_rule_log"`
	BusinessRuleLogUid int32     `bun:"business_rule_log_uid,type:int,pk,identity"`
	UserId             string    `bun:"user_id,type:varchar(255)"`
	LogAction          string    `bun:"log_action,type:varchar(255)"`
	RuleName           string    `bun:"rule_name,type:varchar(255),nullzero"`
	RuleManagerName    string    `bun:"rule_manager_name,type:varchar(255),nullzero"`
	RuleAssemblyName   string    `bun:"rule_assembly_name,type:varchar(255),nullzero"`
	RunType            string    `bun:"run_type,type:varchar(255),nullzero"`
	Xml                string    `bun:"xml,type:text(2147483647),nullzero"`
	ReturnValue        string    `bun:"return_value,type:varchar(255),nullzero"`
	ReturnMessage      string    `bun:"return_message,type:varchar(8000),nullzero"`
	UpdateClassName    string    `bun:"update_class_name,type:varchar(255),nullzero"`
	UpdateRowId        int32     `bun:"update_row_id,type:int,nullzero"`
	UpdateRowNumber    int32     `bun:"update_row_number,type:int,nullzero"`
	UpdateFieldName    string    `bun:"update_field_name,type:varchar(255),nullzero"`
	UpdateFieldValue   string    `bun:"update_field_value,type:varchar(8000),nullzero"`
	DateCreated        time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy          string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified   time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy   string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
