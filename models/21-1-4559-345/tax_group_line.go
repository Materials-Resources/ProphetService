package model

type TaxGroupLine struct {
	bun.BaseModel    `bun:"table:tax_group_line"`
	TaxGroupId       string    `bun:"tax_group_id,type:varchar(10),pk"`
	CompanyId        string    `bun:"company_id,type:varchar(8),pk"`
	JurisdictionId   string    `bun:"jurisdiction_id,type:varchar(10),pk"`
	DeleteFlag       string    `bun:"delete_flag,type:char"`
	DateCreated      time.Time `bun:"date_created,type:datetime"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"`
	Taxable          string    `bun:"taxable,type:char,nullzero"`
}
