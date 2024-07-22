package prophet

import "github.com/uptrace/bun"

type Counter struct {
	bun.BaseModel `bun:"table:counter"`
	Id            string  `bun:"id,type:varchar(64),pk"`        // What is the unique identifier for this counter?
	Description   *string `bun:"description,type:varchar(255)"` // How would you describe this repeating journal entry?
	CounterNum    int32   `bun:"counter_num,type:int"`          // Numerical value of counter_no.
	SeqName       *string `bun:"seq_name,type:Error: 50000, Severity: -1, State: 1. (Params:). The error is printed in terse mode because there was error during formatting. Tracing, ETW, notifications etc are skipped.\r\n"`
	TableName     *string `bun:"table_name,type:Error: 50000, Severity: -1, State: 1. (Params:). The error is printed in terse mode because there was error during formatting. Tracing, ETW, notifications etc are skipped.\r\n"`
	ColumnName    *string `bun:"column_name,type:Error: 50000, Severity: -1, State: 1. (Params:). The error is printed in terse mode because there was error during formatting. Tracing, ETW, notifications etc are skipped.\r\n"`
}
