package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type CopyTableDataXArgument struct {
	bun.BaseModel             `bun:"table:copy_table_data_x_argument"`
	CopyTableDataXArgumentUid int32     `bun:"copy_table_data_x_argument_uid,type:int,autoincrement,identity,pk"` // Unique identifier of table (Identity Column)
	CopyTableDataXProcessUid  int32     `bun:"copy_table_data_x_process_uid,type:int,unique"`                     // Process Type Code
	ArgumentName              string    `bun:"argument_name,type:varchar(255),unique"`                            // Argument Name
	DefaultValue              string    `bun:"default_value,type:varchar(255)"`                                   // Argument Default Value
	DateCreated               time.Time `bun:"date_created,type:datetime,default:(getdate())"`                    // Date and time the record was originally created
	CreatedBy                 string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`              // User who created the record
	DateLastModified          time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`              // Date and time the record was modified
	LastMaintainedBy          string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`      // User who last changed the record
}
