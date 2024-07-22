package prophet

import "github.com/uptrace/bun"

type SupportSqlColumnDetail struct {
	bun.BaseModel             `bun:"table:support_sql_column_detail"`
	SupportSqlColumnDetailUid int32  `bun:"support_sql_column_detail_uid,type:int,autoincrement,unique,identity"`
	SupportQueryCd            int32  `bun:"support_query_cd,type:int,pk"`
	ColumnSequence            int32  `bun:"column_sequence,type:int,pk"`
	ColumnLabel               string `bun:"column_label,type:varchar(255)"`
	ColumnToken               string `bun:"column_token,type:varchar(255)"`
	ColumnDatatypeCd          int32  `bun:"column_datatype_cd,type:int"`
	ColumnLength              *int32 `bun:"column_length,type:int"`
	ColumnPrecision           *int32 `bun:"column_precision,type:int"`
}
