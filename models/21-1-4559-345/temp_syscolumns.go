package model

import "github.com/uptrace/bun"

type TempSyscolumns struct {
	bun.BaseModel `bun:"table:temp_syscolumns"`
	ColumnName    string `bun:"column_name,type:nvarchar(128),nullzero"`
	TableName     string `bun:"table_name,type:nvarchar(128)"`
}
