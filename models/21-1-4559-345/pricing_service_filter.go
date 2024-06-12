package model

type PricingServiceFilter struct {
	bun.BaseModel    `bun:"table:pricing_service_filter"`
	LayoutId         float64   `bun:"layout_id,type:decimal(19,0),pk"`
	LayoutDetailId   float64   `bun:"layout_detail_id,type:decimal(19,0),pk"`
	FilterId         float64   `bun:"filter_id,type:decimal(19,0),pk"`
	Operand          string    `bun:"operand,type:varchar(255),nullzero"`
	FilterValue      string    `bun:"filter_value,type:varchar(50),nullzero"`
	DeleteFlag       string    `bun:"delete_flag,type:char"`
	DateCreated      time.Time `bun:"date_created,type:datetime"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
}
