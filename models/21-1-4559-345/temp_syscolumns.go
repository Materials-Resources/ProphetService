package model

type TempSyscolumns struct {
	bun.BaseModel `bun:"table:temp_syscolumns"`
	ColumnName    `bun:"column_name,type:nvarchar(128),nullzero"`
	TableName     `bun:"table_name,type:nvarchar(128)"`
}
