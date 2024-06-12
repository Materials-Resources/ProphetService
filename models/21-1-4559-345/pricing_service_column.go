package model

type PricingServiceColumn struct {
	bun.BaseModel           `bun:"table:pricing_service_column"`
	ColumnId                float64   `bun:"column_id,type:decimal(19,0),pk"`
	ColumnName              string    `bun:"column_name,type:varchar(40)"`
	ColumnDesc              string    `bun:"column_desc,type:varchar(255),nullzero"`
	AlternateColumnName     string    `bun:"alternate_column_name,type:varchar(40),nullzero"`
	InventoryDefaultsName   string    `bun:"inventory_defaults_name,type:varchar(40),nullzero"`
	RequiredInd             string    `bun:"required_ind,type:char,nullzero"`
	DefaultInd              string    `bun:"default_ind,type:char,nullzero"`
	DeleteFlag              string    `bun:"delete_flag,type:char"`
	DateCreated             time.Time `bun:"date_created,type:datetime"`
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"`
	DataType                string    `bun:"data_type,type:varchar(10),nullzero"`
	ColumnSize              float64   `bun:"column_size,type:decimal(19,0),nullzero"`
	ValidationRule          string    `bun:"validation_rule,type:char,default:('N')"`
	TpcxColumnSequence      int32     `bun:"tpcx_column_sequence,type:int,default:(0)"`
	TpcxOnly                string    `bun:"tpcx_only,type:char,default:('N')"`
	TpcxAbbrevColumnSeq     int32     `bun:"tpcx_abbrev_column_seq,type:int,default:(0)"`
	TpcxAbbrevColumnName    string    `bun:"tpcx_abbrev_column_name,type:varchar(40),nullzero"`
	TpcxAbbrevDate          time.Time `bun:"tpcx_abbrev_date,type:datetime,nullzero"`
	TpcxAbbrevAddForUpdate  string    `bun:"tpcx_abbrev_add_for_update,type:char,default:('N')"`
	ActLayoutOnly           string    `bun:"act_layout_only,type:char,default:('N')"`
	ActLayoutColumnSeq      int32     `bun:"act_layout_column_seq,type:int,default:(0)"`
	ActLayoutColumnName     string    `bun:"act_layout_column_name,type:varchar(40),nullzero"`
	ActLayoutDate           time.Time `bun:"act_layout_date,type:datetime,nullzero"`
	ActLayoutAddForUpdate   string    `bun:"act_layout_add_for_update,type:char,default:('N')"`
	AqnetLayoutOnly         string    `bun:"aqnet_layout_only,type:char,default:('N')"`
	AqnetLayoutColumnSeq    int32     `bun:"aqnet_layout_column_seq,type:int,default:((0))"`
	AqnetLayoutColumnName   string    `bun:"aqnet_layout_column_name,type:varchar(40),nullzero"`
	AqnetLayoutDate         time.Time `bun:"aqnet_layout_date,type:datetime,nullzero"`
	AqnetLayoutAddForUpdate string    `bun:"aqnet_layout_add_for_update,type:char,default:('N')"`
	PricingTemplateDfltType string    `bun:"pricing_template_dflt_type,type:char,default:('N')"`
	ImportNewValuesFlag     string    `bun:"import_new_values_flag,type:char,default:('N')"`
}
