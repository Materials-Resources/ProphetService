package model

type ImportSuspenseSettings struct {
	bun.BaseModel    `bun:"table:import_suspense_settings"`
	ImpexpSourceId   string    `bun:"impexp_source_id,type:varchar(50),pk"`
	TransactionSetId string    `bun:"transaction_set_id,type:varchar(32),pk"`
	RowStatusFlag    int32     `bun:"row_status_flag,type:int,default:(705)"`
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
}
