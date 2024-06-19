package gen

import "github.com/uptrace/bun"

type P21DeletionSql struct {
	bun.BaseModel     `bun:"table:p21_deletion_sql"`
	P21DeletionSqlUid int32  `bun:"p21_deletion_sql_uid,type:int,autoincrement,scanonly"`
	DeletionType      int32  `bun:"deletion_type,type:int"`
	DeletionOrder     int32  `bun:"deletion_order,type:int"`
	DeletionDb        string `bun:"deletion_db,type:char(1)"`
	DeletionAction    string `bun:"deletion_action,type:char(1)"`
	DeletionSql       string `bun:"deletion_sql,type:Error: 50000, Severity: -1, State: 1. (Params:). The error is printed in terse mode because there was error during formatting. Tracing, ETW, notifications etc are skipped.\r\n"`
	DeletionTable     string `bun:"deletion_table,type:varchar(256)"`
}
