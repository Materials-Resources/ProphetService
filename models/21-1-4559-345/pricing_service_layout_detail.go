package model

type PricingServiceLayoutDetail struct {
	bun.BaseModel       `bun:"table:pricing_service_layout_detail"`
	LayoutId            float64   `bun:"layout_id,type:decimal(19,0),pk"`
	LayoutDetailId      float64   `bun:"layout_detail_id,type:decimal(19,0),pk"`
	ColumnName          string    `bun:"column_name,type:varchar(50),nullzero"`
	StartPosition       float64   `bun:"start_position,type:decimal(19,0),nullzero"`
	DataPrecision       float64   `bun:"data_precision,type:decimal(19,0),nullzero"`
	DataType            string    `bun:"data_type,type:varchar(10),nullzero"`
	DecimalScale        float64   `bun:"decimal_scale,type:decimal(19,0),nullzero"`
	AcsType             string    `bun:"acs_type,type:varchar(10),nullzero"`
	ExtensionInd        string    `bun:"extension_ind,type:char,nullzero"`
	ValueInd            string    `bun:"value_ind,type:char,nullzero"`
	ValueId             float64   `bun:"value_id,type:decimal(19,0),nullzero"`
	ConversionInd       string    `bun:"conversion_ind,type:char,nullzero"`
	FilterInd           string    `bun:"filter_ind,type:char,nullzero"`
	DeleteFlag          string    `bun:"delete_flag,type:char"`
	DateCreated         time.Time `bun:"date_created,type:datetime"`
	DateLastModified    time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy    string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	DataLength          float64   `bun:"data_length,type:decimal(19,0),nullzero"`
	RequiredInd         string    `bun:"required_ind,type:char,nullzero"`
	DefaultInd          string    `bun:"default_ind,type:char,nullzero"`
	NewItemsOnly        string    `bun:"new_items_only,type:char,nullzero"`
	DateMask            string    `bun:"date_mask,type:varchar(255),nullzero"`
	StartPositionDate   time.Time `bun:"start_position_date,type:datetime,default:(getdate())"`
	SetByOption         string    `bun:"set_by_option,type:char,nullzero"`
	ImportNewValuesFlag string    `bun:"import_new_values_flag,type:char,nullzero"`
	ItemCatalogDefUid   int32     `bun:"item_catalog_def_uid,type:int,nullzero"`
}
