package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type DctTransactionTable struct {
	bun.BaseModel          `bun:"table:dct_transaction_table"`
	DctTransactionTableUid int32     `bun:"dct_transaction_table_uid,type:int,pk"`                        // unique identifier
	DctTransactionUid      int32     `bun:"dct_transaction_uid,type:int"`                                 // foreign key to dct_transaction table
	SequenceNo             int32     `bun:"sequence_no,type:int"`                                         // process order of tables within transaction
	TableId                string    `bun:"table_id,type:varchar(255)"`                                   // table to be updated during the import
	TableDesc              string    `bun:"table_desc,type:varchar(255)"`                                 // description of table to be updated during import
	LegacyFileSuffix       *string   `bun:"legacy_file_suffix,type:varchar(255)"`                         // code to identify legacy export file for this table
	BusinessObjectName     *string   `bun:"business_object_name,type:varchar(255)"`                       // business object that will handle validating and updating this table
	SetupEventName         *string   `bun:"setup_event_name,type:varchar(255)"`                           // event to be called to setup business object
	BaselineTableFlag      string    `bun:"baseline_table_flag,type:char(1),default:('Y')"`               // indicates whether table is baseline or custom
	RowStatusFlag          int32     `bun:"row_status_flag,type:int,default:(704)"`                       // status
	DateCreated            time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy              string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified       time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy       string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
