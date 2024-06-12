package model

import (
	"github.com/uptrace/bun"
	"time"
)

type ImportSuspenseHdr struct {
	bun.BaseModel        `bun:"table:import_suspense_hdr"`
	ImportSuspenseHdrUid int32     `bun:"import_suspense_hdr_uid,type:int,pk"`
	ImpexpSourceId       string    `bun:"impexp_source_id,type:varchar(50)"`
	TransactionSetId     string    `bun:"transaction_set_id,type:varchar(32)"`
	MasterFilePath       string    `bun:"master_file_path,type:varchar(255)"`
	MasterFileName       string    `bun:"master_file_name,type:varchar(255)"`
	RequiredForImport    string    `bun:"required_for_import,type:varchar(255)"`
	DateCreated          time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	DateLastModified     time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy     string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	ScheduledImport      string    `bun:"scheduled_import,type:char,default:('N')"`
	ImpexpSourceDesc     string    `bun:"impexp_source_desc,type:varchar(255),nullzero"`
	TransactionSetDesc   string    `bun:"transaction_set_desc,type:varchar(255),nullzero"`
	FileFormatCd         int32     `bun:"file_format_cd,type:int,default:(1012)"`
	TransactionSusPath   string    `bun:"transaction_sus_path,type:varchar(200),nullzero"`
}
