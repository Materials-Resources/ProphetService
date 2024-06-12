package model

type PricingServiceExtension struct {
	bun.BaseModel        `bun:"table:pricing_service_extension"`
	LayoutId             float64   `bun:"layout_id,type:decimal(19,0),pk"`
	LayoutDetailId       float64   `bun:"layout_detail_id,type:decimal(19,0),pk"`
	ExtensionId          float64   `bun:"extension_id,type:decimal(19,0),pk"`
	ExtensionSource      string    `bun:"extension_source,type:varchar(10),nullzero"`
	ExtensionValue       string    `bun:"extension_value,type:varchar(50),nullzero"`
	SourceLayoutId       float64   `bun:"source_layout_id,type:decimal(19,0),nullzero"`
	SourceLayoutDetailId float64   `bun:"source_layout_detail_id,type:decimal(19,0),nullzero"`
	SequenceNumber       float64   `bun:"sequence_number,type:decimal(19,0),nullzero"`
	DeleteFlag           string    `bun:"delete_flag,type:char"`
	DateCreated          time.Time `bun:"date_created,type:datetime"`
	DateLastModified     time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy     string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	ExtensionPosition    string    `bun:"extension_position,type:varchar(10),nullzero"`
	SourceStart          int32     `bun:"source_start,type:int,nullzero"`
	SourceLength         int32     `bun:"source_length,type:int,nullzero"`
}
