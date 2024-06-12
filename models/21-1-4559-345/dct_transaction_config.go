package model

import (
	"github.com/uptrace/bun"
	"time"
)

type DctTransactionConfig struct {
	bun.BaseModel           `bun:"table:dct_transaction_config"`
	DctTransactionConfigUid int32     `bun:"dct_transaction_config_uid,type:int,pk"`
	DctTransactionUid       int32     `bun:"dct_transaction_uid,type:int"`
	ConfigurationId         int32     `bun:"configuration_id,type:int"`
	RowStatusFlag           int32     `bun:"row_status_flag,type:int,default:(704)"`
	DateCreated             time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy               string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
