package model

type PendingImport struct {
	bun.BaseModel         `bun:"table:pending_import"`
	PendingImportUid      int32     `bun:"pending_import_uid,type:int,pk"`
	ImpexpSourceId        string    `bun:"impexp_source_id,type:varchar(50)"`
	TransactionSetId      string    `bun:"transaction_set_id,type:varchar(32)"`
	DbtableId             string    `bun:"dbtable_id,type:varchar(40)"`
	PendingImportSetNo    int32     `bun:"pending_import_set_no,type:int"`
	ImportData            string    `bun:"import_data,type:text(2147483647)"`
	StatusFlag            int32     `bun:"status_flag,type:int"`
	DateCreated           time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy             string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified      time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy      string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	ImportedTransactionId int32     `bun:"imported_transaction_id,type:int,nullzero"`
	SourceReferenceNo     string    `bun:"source_reference_no,type:varchar(255),nullzero"`
	ImportResultInfo      string    `bun:"import_result_info,type:varchar(8000),nullzero"`
}
