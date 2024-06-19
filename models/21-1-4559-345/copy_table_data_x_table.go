package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type CopyTableDataXTable struct {
	bun.BaseModel            `bun:"table:copy_table_data_x_table"`
	CopyTableDataXTableUid   int32     `bun:"copy_table_data_x_table_uid,type:int,autoincrement,scanonly,pk"` // Unique identifier of table (Identity Column)
	CopyTableDataXProcessUid int32     `bun:"copy_table_data_x_process_uid,type:int,unique"`                  // Copy Table Process
	TableName                string    `bun:"table_name,type:varchar(255),unique"`                            // Table Name
	SourceTableName          string    `bun:"source_table_name,type:varchar(255),nullzero"`                   // Source Table Name
	SequenceNumber           int32     `bun:"sequence_number,type:int"`                                       // Sequence Order Number
	DateCreated              time.Time `bun:"date_created,type:datetime,default:(getdate())"`                 // Date and time the record was originally created
	CreatedBy                string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`           // User who created the record
	DateLastModified         time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`           // Date and time the record was modified
	LastMaintainedBy         string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`   // User who last changed the record
	UseReadUncommitted       string    `bun:"use_read_uncommitted,type:char(1),default:('N')"`                // Whether the source table will use dirty reads for the select
}
