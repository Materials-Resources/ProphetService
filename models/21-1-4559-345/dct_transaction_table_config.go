package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type DctTransactionTableConfig struct {
	bun.BaseModel                `bun:"table:dct_transaction_table_config"`
	DctTransactionTableConfigUid int32     `bun:"dct_transaction_table_config_uid,type:int,pk"`                 // unique identifier
	DctTransactionTableUid       int32     `bun:"dct_transaction_table_uid,type:int"`                           // foreign key to dct_transaction_table
	ConfigurationId              int32     `bun:"configuration_id,type:int"`                                    // configuration to which the associated table within transaction applies
	RowStatusFlag                int32     `bun:"row_status_flag,type:int,default:(704)"`                       // status
	DateCreated                  time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy                    string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified             time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy             string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
