package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type CopyTableDataXProcess struct {
	bun.BaseModel            `bun:"table:copy_table_data_x_process"`
	CopyTableDataXProcessUid int32     `bun:"copy_table_data_x_process_uid,type:int,autoincrement,pk"`      // Unique identifier of table (Identity Column)
	ProcessTypeCd            int32     `bun:"process_type_cd,type:int"`                                     // Process Type Code
	DateCreated              time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy                string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified         time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy         string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
