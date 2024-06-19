package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type CopyTableDataXColumn struct {
	bun.BaseModel           `bun:"table:copy_table_data_x_column"`
	CopyTableDataXColumnUid int32     `bun:"copy_table_data_x_column_uid,type:int,autoincrement,pk"`       // Unique identifier of table (Identity Column)
	CopyTableDataXTableUid  int32     `bun:"copy_table_data_x_table_uid,type:int,unique"`                  // Table associated with Columns
	ColumnName              string    `bun:"column_name,type:varchar(255),unique"`                         // Column Name
	ColumnValue             string    `bun:"column_value,type:varchar(255),nullzero"`                      // Value to replace the Column Name
	RecordTypeCd            int32     `bun:"record_type_cd,type:int,unique"`                               // Defines the type of Record (Static or Dynamic)
	CodeGroupCd             int16     `bun:"code_group_cd,type:smallint,nullzero"`                         // Codes values that can replace the column name (Only used by Dynamic records)
	DefaultValueCd          int32     `bun:"default_value_cd,type:int,nullzero"`                           // Default Value for column  (Only used by Dynamic records)
	ColumnAliasName         string    `bun:"column_alias_name,type:varchar(255),nullzero"`                 // Alias for Column
	ColumnDisplayName       string    `bun:"column_display_name,type:varchar(255),nullzero"`               // Display Name for column  (Only used by Dynamic records)
	DateCreated             time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy               string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
