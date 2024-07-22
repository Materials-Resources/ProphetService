package prophet

import "github.com/uptrace/bun"

type EccGetTableSchemaTable struct {
	bun.BaseModel             `bun:"table:ecc_get_table_schema_table"`
	EccGetTableSchemaTableUid int32  `bun:"ecc_get_table_schema_table_uid,type:int,autoincrement,identity"`
	TableName                 string `bun:"table_name,type:Error: 50000, Severity: -1, State: 1. (Params:). The error is printed in terse mode because there was error during formatting. Tracing, ETW, notifications etc are skipped.\r\n"`
}
