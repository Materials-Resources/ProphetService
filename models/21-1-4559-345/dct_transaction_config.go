package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type DctTransactionConfig struct {
	bun.BaseModel           `bun:"table:dct_transaction_config"`
	DctTransactionConfigUid int32     `bun:"dct_transaction_config_uid,type:int,pk"`                       // unique identifier
	DctTransactionUid       int32     `bun:"dct_transaction_uid,type:int"`                                 // foreign key to dct_transaction table
	ConfigurationId         int32     `bun:"configuration_id,type:int"`                                    // database system_setting configuration ID
	RowStatusFlag           int32     `bun:"row_status_flag,type:int,default:(704)"`                       // record status
	DateCreated             time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy               string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
