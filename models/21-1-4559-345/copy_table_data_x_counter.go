package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type CopyTableDataXCounter struct {
	bun.BaseModel            `bun:"table:copy_table_data_x_counter"`
	CopyTableDataXCounterUid int32     `bun:"copy_table_data_x_counter_uid,type:int,autoincrement,identity,pk"` // Unique identifier of table (Identity Column)
	CopyTableDataXTableUid   int32     `bun:"copy_table_data_x_table_uid,type:int,unique"`                      // Table associated with Counter
	CounterId                string    `bun:"counter_id,type:varchar(64),unique"`                               // Counter associated with table
	DateCreated              time.Time `bun:"date_created,type:datetime,default:(getdate())"`                   // Date and time the record was originally created
	CreatedBy                string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`             // User who created the record
	DateLastModified         time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`             // Date and time the record was modified
	LastMaintainedBy         string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`     // User who last changed the record
}
