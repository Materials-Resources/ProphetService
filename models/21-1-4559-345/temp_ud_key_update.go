package gen

import "github.com/uptrace/bun"

type TempUdKeyUpdate struct {
	bun.BaseModel      `bun:"table:temp_ud_key_update"`
	TempUdKeyUpdateUid int32  `bun:"temp_ud_key_update_uid,type:int,autoincrement,identity"`
	TableName          string `bun:"table_name,type:Error: 50000, Severity: -1, State: 1. (Params:). The error is printed in terse mode because there was error during formatting. Tracing, ETW, notifications etc are skipped.\r\n"`
	BaseTable          string `bun:"base_table,type:Error: 50000, Severity: -1, State: 1. (Params:). The error is printed in terse mode because there was error during formatting. Tracing, ETW, notifications etc are skipped.\r\n"`
	UdKeyName          string `bun:"ud_key_name,type:Error: 50000, Severity: -1, State: 1. (Params:). The error is printed in terse mode because there was error during formatting. Tracing, ETW, notifications etc are skipped.\r\n"`
}
