package prophet

import "github.com/uptrace/bun"

type CcColumns struct {
	bun.BaseModel     `bun:"table:cc_columns"`
	ColumnUid         int32   `bun:"column_uid,type:int,pk"`
	TableUid          int32   `bun:"table_uid,type:int"`
	ColumnName        string  `bun:"column_name,type:char(255)"`
	ColumnDescription *string `bun:"column_description,type:char(255)"`
}
