package model

import (
	"github.com/uptrace/bun"
	"time"
)

type UserDefinedColumn struct {
	bun.BaseModel        `bun:"table:user_defined_column"`
	UserDefinedColumnUid int32     `bun:"user_defined_column_uid,type:int,pk,identity"`
	BaseTable            string    `bun:"base_table,type:varchar(40)"`
	ColumnName           string    `bun:"column_name,type:varchar(40)"`
	DataType             string    `bun:"data_type,type:varchar(10)"`
	DataLength           int32     `bun:"data_length,type:int"`
	DataScale            int32     `bun:"data_scale,type:int,nullzero"`
	ColumnDescription    string    `bun:"column_description,type:varchar(255)"`
	ColumnLabel          string    `bun:"column_label,type:varchar(255)"`
	NewTableName         string    `bun:"new_table_name,type:varchar(255)"`
	DateCreated          time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy            string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified     time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy     string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
