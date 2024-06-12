package model

type JurisdictionAcct struct {
	bun.BaseModel    `bun:"table:jurisdiction_acct"`
	CompanyNo        string    `bun:"company_no,type:varchar(8),pk"`
	JurisdictionId   string    `bun:"jurisdiction_id,type:varchar(10),pk"`
	GlAccountNo      string    `bun:"gl_account_no,type:varchar(32)"`
	DeleteFlag       string    `bun:"delete_flag,type:char"`
	DateCreated      time.Time `bun:"date_created,type:datetime"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"`
}
