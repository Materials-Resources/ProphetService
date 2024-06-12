package model

type PricingServiceLayout struct {
	bun.BaseModel           `bun:"table:pricing_service_layout"`
	LayoutId                float64   `bun:"layout_id,type:decimal(19,0),pk"`
	LayoutType              string    `bun:"layout_type,type:varchar(10),nullzero"`
	LayoutDesc              string    `bun:"layout_desc,type:varchar(50),nullzero"`
	LayoutKey               string    `bun:"layout_key,type:varchar(10),nullzero"`
	LayoutPath              string    `bun:"layout_path,type:varchar(255),nullzero"`
	LayoutFilename          string    `bun:"layout_filename,type:varchar(50),nullzero"`
	DeleteFlag              string    `bun:"delete_flag,type:char"`
	DateCreated             time.Time `bun:"date_created,type:datetime"`
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	LayoutTabletype         string    `bun:"layout_tabletype,type:varchar(10),nullzero"`
	LayoutDelimiter         string    `bun:"layout_delimiter,type:char,nullzero"`
	LayoutFirstlinenames    string    `bun:"layout_firstlinenames,type:char,nullzero"`
	CompanyId               string    `bun:"company_id,type:varchar(8)"`
	LocationId              float64   `bun:"location_id,type:decimal(19,0)"`
	AddLocation             string    `bun:"add_location,type:char,nullzero"`
	AddSupplier             string    `bun:"add_supplier,type:char,nullzero"`
	TpcxFormatCd            int32     `bun:"tpcx_format_cd,type:int,default:(0)"`
	PurchaseTransferGroupId string    `bun:"purchase_transfer_group_id,type:varchar(8),nullzero"`
	PricingTemplateId       string    `bun:"pricing_template_id,type:varchar(255),nullzero"`
	CatalogStatusFlag       string    `bun:"catalog_status_flag,type:char,default:('N')"`
	DataSourceCd            int32     `bun:"data_source_cd,type:int,default:((2564))"`
	DataDestinationCd       int32     `bun:"data_destination_cd,type:int,default:((2566))"`
	BuildInBothFlag         string    `bun:"build_in_both_flag,type:char,default:('N')"`
	ApplyValueListFirstFlag string    `bun:"apply_value_list_first_flag,type:char,default:('Y')"`
	TempDataStorageCd       int32     `bun:"temp_data_storage_cd,type:int,default:((2570))"`
}
