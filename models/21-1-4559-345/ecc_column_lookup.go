package gen

import "github.com/uptrace/bun"

type EccColumnLookup struct {
	bun.BaseModel      `bun:"table:ecc_column_lookup"`
	EccColumnLookupUid int32  `bun:"ecc_column_lookup_uid,type:int,autoincrement,identity,pk"`
	TransactionType    string `bun:"transaction_type,type:Error: 50000, Severity: -1, State: 1. (Params:). The error is printed in terse mode because there was error during formatting. Tracing, ETW, notifications etc are skipped.\r\n"`
	ParameterTag       string `bun:"parameter_tag,type:Error: 50000, Severity: -1, State: 1. (Params:). The error is printed in terse mode because there was error during formatting. Tracing, ETW, notifications etc are skipped.\r\n"`
	TableName          string `bun:"table_name,type:Error: 50000, Severity: -1, State: 1. (Params:). The error is printed in terse mode because there was error during formatting. Tracing, ETW, notifications etc are skipped.\r\n"`
	ColumnName         string `bun:"column_name,type:Error: 50000, Severity: -1, State: 1. (Params:). The error is printed in terse mode because there was error during formatting. Tracing, ETW, notifications etc are skipped.\r\n"`
	Datatype           string `bun:"datatype,type:Error: 50000, Severity: -1, State: 1. (Params:). The error is printed in terse mode because there was error during formatting. Tracing, ETW, notifications etc are skipped.\r\n"`
}
