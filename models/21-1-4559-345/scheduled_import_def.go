package model

import (
	"github.com/uptrace/bun"
	"time"
)

type ScheduledImportDef struct {
	bun.BaseModel         `bun:"table:scheduled_import_def"`
	ScheduledImportDefUid int32     `bun:"scheduled_import_def_uid,type:int,pk"`
	ImpexpSourceUid       int32     `bun:"impexp_source_uid,type:int"`
	TransactionSetUid     int32     `bun:"transaction_set_uid,type:int"`
	DbtableId             string    `bun:"dbtable_id,type:varchar(40)"`
	DbtableDesc           string    `bun:"dbtable_desc,type:varchar(40)"`
	Req                   string    `bun:"req,type:char"`
	DateCreated           time.Time `bun:"date_created,type:datetime"`
	DateLastModified      time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy      string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	FileFormatCd          int32     `bun:"file_format_cd,type:int,default:(1012)"`
	DefaultFileprefix     string    `bun:"default_fileprefix,type:varchar(255),nullzero"`
}
