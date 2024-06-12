package model

type ScheduledImportMaster struct {
	bun.BaseModel            `bun:"table:scheduled_import_master"`
	ScheduledImportMasterUid int32     `bun:"scheduled_import_master_uid,type:int,pk"`
	ImpexpSourceUid          int32     `bun:"impexp_source_uid,type:int"`
	TransactionSetUid        int32     `bun:"transaction_set_uid,type:int"`
	PollingPath              string    `bun:"polling_path,type:varchar(200)"`
	TransactionLogPath       string    `bun:"transaction_log_path,type:varchar(200)"`
	TransactionSumPath       string    `bun:"transaction_sum_path,type:varchar(200)"`
	TransactionSusPath       string    `bun:"transaction_sus_path,type:varchar(200)"`
	TransactionErrPath       string    `bun:"transaction_err_path,type:varchar(200)"`
	Active                   string    `bun:"active,type:char"`
	DateCreated              time.Time `bun:"date_created,type:datetime"`
	DateLastModified         time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy         string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	FileFormatCd             int32     `bun:"file_format_cd,type:int,default:(1012)"`
	XmlDocumentUid           int32     `bun:"xml_document_uid,type:int,nullzero"`
	FileLockingFlag          string    `bun:"file_locking_flag,type:char,default:('N')"`
}
