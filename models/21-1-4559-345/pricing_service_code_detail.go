package model

type PricingServiceCodeDetail struct {
	bun.BaseModel    `bun:"table:pricing_service_code_detail"`
	CodeId           float64   `bun:"code_id,type:decimal(19,0),pk"`
	CodeDetailId     float64   `bun:"code_detail_id,type:decimal(19,0),pk"`
	CodeValue        string    `bun:"code_value,type:varchar(40),nullzero"`
	ValueDesc        string    `bun:"value_desc,type:varchar(254),nullzero"`
	GroupCode        string    `bun:"group_code,type:varchar(10),nullzero"`
	SubgroupCode     string    `bun:"subgroup_code,type:varchar(10),nullzero"`
	DeleteFlag       string    `bun:"delete_flag,type:char"`
	DateCreated      time.Time `bun:"date_created,type:datetime"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"`
}
