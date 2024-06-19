package gen

import "github.com/uptrace/bun"

type TempSyscolumns struct {
	bun.BaseModel `bun:"table:temp_syscolumns"`
	ColumnName    string `bun:"column_name,type:Error: 50000, Severity: -1, State: 1. (Params:). The error is printed in terse mode because there was error during formatting. Tracing, ETW, notifications etc are skipped.\r\n,nullzero"`
	TableName     string `bun:"table_name,type:sysname"`
}
