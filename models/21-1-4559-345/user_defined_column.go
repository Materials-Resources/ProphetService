package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type UserDefinedColumn struct {
	bun.BaseModel        `bun:"table:user_defined_column"`
	UserDefinedColumnUid int32     `bun:"user_defined_column_uid,type:int,autoincrement,pk"` // Unique Identifier for the user_defined_column record
	BaseTable            string    `bun:"base_table,type:varchar(40)"`                       // The base table for the user defined column
	ColumnName           string    `bun:"column_name,type:varchar(40)"`                      // The name of the user defined column
	DataType             string    `bun:"data_type,type:varchar(10)"`                        // The data type of the user defined column
	DataLength           int32     `bun:"data_length,type:int"`                              // The data length of the user defined column
	DataScale            int32     `bun:"data_scale,type:int,nullzero"`                      // The data scale of the user defined column
	ColumnDescription    string    `bun:"column_description,type:varchar(255)"`              // The description of the the user defined column
	ColumnLabel          string    `bun:"column_label,type:varchar(255)"`                    // The label for the user defined column
	NewTableName         string    `bun:"new_table_name,type:varchar(255)"`
	DateCreated          time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy            string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified     time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy     string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
