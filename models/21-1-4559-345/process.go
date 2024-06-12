package model

type Process struct {
	bun.BaseModel           `bun:"table:process"`
	ProcessUid              int32     `bun:"process_uid,type:int,pk"`
	ProcessCode             string    `bun:"process_code,type:varchar(255)"`
	ProcessName             string    `bun:"process_name,type:varchar(255)"`
	ProcessDescription      string    `bun:"process_description,type:varchar(255),nullzero"`
	DateCreated             time.Time `bun:"date_created,type:datetime"`
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"`
	RowStatusFlag           int16     `bun:"row_status_flag,type:smallint"`
	CompanyId               string    `bun:"company_id,type:varchar(8),nullzero"`
	LocationId              float64   `bun:"location_id,type:decimal(19,0),nullzero"`
	AllLocationsFlag        string    `bun:"all_locations_flag,type:char,default:('Y')"`
	ReceiptProcessFlag      string    `bun:"receipt_process_flag,type:char,default:('N')"`
	RawInvMastUid           int32     `bun:"raw_inv_mast_uid,type:int,nullzero"`
	FinishedInvMastUid      int32     `bun:"finished_inv_mast_uid,type:int,nullzero"`
	RawYieldQty             float64   `bun:"raw_yield_qty,type:decimal(19,9),nullzero"`
	FinishedYieldQty        float64   `bun:"finished_yield_qty,type:decimal(19,9),nullzero"`
	RawYieldUom             string    `bun:"raw_yield_uom,type:varchar(8),nullzero"`
	FinishedYieldUom        string    `bun:"finished_yield_uom,type:varchar(8),nullzero"`
	CaptureUsageFlag        string    `bun:"capture_usage_flag,type:char,default:('N')"`
	RawItemRevisionUid      int32     `bun:"raw_item_revision_uid,type:int,nullzero"`
	FinishedItemRevisionUid int32     `bun:"finished_item_revision_uid,type:int,nullzero"`
	ProductionOrderFlag     string    `bun:"production_order_flag,type:char,nullzero"`
	OneOffFlag              string    `bun:"one_off_flag,type:char,nullzero"`
}
