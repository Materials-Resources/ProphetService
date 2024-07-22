package prophet

import "github.com/uptrace/bun"

type ClassExpansion struct {
	bun.BaseModel       `bun:"table:class_expansion"`
	ClassExpansionUid   int32  `bun:"class_expansion_uid,type:int,autoincrement,identity"`
	SchemaName          string `bun:"schema_name,type:varchar(255)"`
	TableName           string `bun:"table_name,type:varchar(255)"`
	ColumnName          string `bun:"column_name,type:varchar(255)"`
	NullState           string `bun:"null_state,type:varchar(255)"`
	Done                string `bun:"done,type:char(1)"`
	ColumnExpansionTime *int32 `bun:"column_expansion_time,type:int"`
}
