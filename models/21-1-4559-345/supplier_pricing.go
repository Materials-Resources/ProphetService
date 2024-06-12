package model

type SupplierPricing struct {
	bun.BaseModel    `bun:"table:supplier_pricing"`
	CompanyId        string    `bun:"company_id,type:varchar(8),pk"`
	SupplierId       float64   `bun:"supplier_id,type:decimal(19,0),pk"`
	PricingMethod    string    `bun:"pricing_method,type:varchar(40)"`
	SourcePrice      string    `bun:"source_price,type:varchar(40),nullzero"`
	Multiplier       float64   `bun:"multiplier,type:decimal(19,9),nullzero"`
	DeleteFlag       string    `bun:"delete_flag,type:char"`
	DateCreated      time.Time `bun:"date_created,type:datetime"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"`
}
