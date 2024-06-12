package model

type PricingServiceCode struct {
	bun.BaseModel    `bun:"table:pricing_service_code"`
	CodeId           float64   `bun:"code_id,type:decimal(19,0),pk"`
	CodeName         string    `bun:"code_name,type:varchar(10)"`
	CodeDesc         string    `bun:"code_desc,type:varchar(254),nullzero"`
	DeleteFlag       string    `bun:"delete_flag,type:char"`
	DateCreated      time.Time `bun:"date_created,type:datetime"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
}
