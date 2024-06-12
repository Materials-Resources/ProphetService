package model

type PricingServiceValueDetail struct {
	bun.BaseModel    `bun:"table:pricing_service_value_detail"`
	ValueId          float64   `bun:"value_id,type:decimal(19,0),pk"`
	ValueDetailId    float64   `bun:"value_detail_id,type:decimal(19,0),pk"`
	OriginalValue    string    `bun:"original_value,type:varchar(50),nullzero"`
	ConvertedValue   string    `bun:"converted_value,type:varchar(50),nullzero"`
	DeleteFlag       string    `bun:"delete_flag,type:char"`
	DateCreated      time.Time `bun:"date_created,type:datetime"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
}
