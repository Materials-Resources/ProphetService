package model

import (
	"github.com/uptrace/bun"
	"time"
)

type DctTransactionTable struct {
	bun.BaseModel          `bun:"table:dct_transaction_table"`
	DctTransactionTableUid int32     `bun:"dct_transaction_table_uid,type:int,pk"`
	DctTransactionUid      int32     `bun:"dct_transaction_uid,type:int"`
	SequenceNo             int32     `bun:"sequence_no,type:int"`
	TableId                string    `bun:"table_id,type:varchar(255)"`
	TableDesc              string    `bun:"table_desc,type:varchar(255)"`
	LegacyFileSuffix       string    `bun:"legacy_file_suffix,type:varchar(255),nullzero"`
	BusinessObjectName     string    `bun:"business_object_name,type:varchar(255),nullzero"`
	SetupEventName         string    `bun:"setup_event_name,type:varchar(255),nullzero"`
	BaselineTableFlag      string    `bun:"baseline_table_flag,type:char,default:('Y')"`
	RowStatusFlag          int32     `bun:"row_status_flag,type:int,default:(704)"`
	DateCreated            time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy              string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified       time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy       string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
