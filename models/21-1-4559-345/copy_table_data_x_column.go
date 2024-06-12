package model

import (
	"github.com/uptrace/bun"
	"time"
)

type CopyTableDataXColumn struct {
	bun.BaseModel           `bun:"table:copy_table_data_x_column"`
	CopyTableDataXColumnUid int32     `bun:"copy_table_data_x_column_uid,type:int,pk,identity"`
	CopyTableDataXTableUid  int32     `bun:"copy_table_data_x_table_uid,type:int"`
	ColumnName              string    `bun:"column_name,type:varchar(255)"`
	ColumnValue             string    `bun:"column_value,type:varchar(255),nullzero"`
	RecordTypeCd            int32     `bun:"record_type_cd,type:int"`
	CodeGroupCd             int16     `bun:"code_group_cd,type:smallint,nullzero"`
	DefaultValueCd          int32     `bun:"default_value_cd,type:int,nullzero"`
	ColumnAliasName         string    `bun:"column_alias_name,type:varchar(255),nullzero"`
	ColumnDisplayName       string    `bun:"column_display_name,type:varchar(255),nullzero"`
	DateCreated             time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy               string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
