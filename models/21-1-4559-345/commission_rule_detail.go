package model

type CommissionRuleDetail struct {
	bun.BaseModel    `bun:"table:commission_rule_detail"`
	CompanyId        string    `bun:"company_id,type:varchar(8),pk"`
	CommissionRuleId string    `bun:"commission_rule_id,type:varchar(8),pk"`
	BreakNumber      float64   `bun:"break_number,type:decimal(4,0),pk"`
	BreakValue       float64   `bun:"break_value,type:decimal(19,9)"`
	BreakType        string    `bun:"break_type,type:char"`
	Value            float64   `bun:"value,type:decimal(19,9)"`
	Factor           string    `bun:"factor,type:char,nullzero"`
	Source           string    `bun:"source,type:char,nullzero"`
	DeleteFlag       string    `bun:"delete_flag,type:char"`
	DateCreated      time.Time `bun:"date_created,type:datetime"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
}
