package model

type GporDynamicLookAhead struct {
	bun.BaseModel        `bun:"table:gpor_dynamic_look_ahead"`
	GporRunHdrUid        int32     `bun:"gpor_run_hdr_uid,type:int"`
	InvMastUid           int32     `bun:"inv_mast_uid,type:int"`
	LocationId           float64   `bun:"location_id,type:decimal(19,0)"`
	DynamicLookAheadDate time.Time `bun:"dynamic_look_ahead_date,type:datetime"`
	DateCreated          time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy            string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified     time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy     string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	DrpItemFlag          string    `bun:"drp_item_flag,type:char,default:('N')"`
	TransactionTypeCd    int32     `bun:"transaction_type_cd,type:int,default:((1322))"`
}
