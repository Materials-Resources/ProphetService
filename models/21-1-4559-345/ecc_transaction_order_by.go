package prophet

import "github.com/uptrace/bun"

type EccTransactionOrderBy struct {
	bun.BaseModel            `bun:"table:ecc_transaction_order_by"`
	EccTransactionOrderByUid int32  `bun:"ecc_transaction_order_by_uid,type:int,autoincrement,identity,pk"`
	TransactionType          string `bun:"transaction_type,type:Error: 50000, Severity: -1, State: 1. (Params:). The error is printed in terse mode because there was error during formatting. Tracing, ETW, notifications etc are skipped.\r\n"`
	OrderBy                  string `bun:"order_by,type:Error: 50000, Severity: -1, State: 1. (Params:). The error is printed in terse mode because there was error during formatting. Tracing, ETW, notifications etc are skipped.\r\n"`
}
