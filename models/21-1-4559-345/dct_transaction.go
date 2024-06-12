package model

import (
	"github.com/uptrace/bun"
	"time"
)

type DctTransaction struct {
	bun.BaseModel           `bun:"table:dct_transaction"`
	DctTransactionUid       int32     `bun:"dct_transaction_uid,type:int,pk"`
	TransactionId           string    `bun:"transaction_id,type:varchar(255)"`
	TransactionDesc         string    `bun:"transaction_desc,type:varchar(255)"`
	BaselineTransactionFlag string    `bun:"baseline_transaction_flag,type:char,default:('N')"`
	ProcessTypeCd           int32     `bun:"process_type_cd,type:int,default:(1657)"`
	ImportTypeCd            int32     `bun:"import_type_cd,type:int,default:(1649)"`
	RowStatusFlag           int32     `bun:"row_status_flag,type:int,default:(704)"`
	DateCreated             time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy               string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	MenuclickedName         string    `bun:"menuclicked_name,type:varchar(255),nullzero"`
	LegacyExportFlag        string    `bun:"legacy_export_flag,type:char,default:('N')"`
}
