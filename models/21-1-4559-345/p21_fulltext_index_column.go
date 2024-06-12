package model

import "github.com/uptrace/bun"

type P21FulltextIndexColumn struct {
	bun.BaseModel            `bun:"table:p21_fulltext_index_column"`
	P21FulltextIndexTableUid int32  `bun:"p21_fulltext_index_table_uid,type:int,pk"`
	ColumnName               string `bun:"column_name,type:varchar(256),pk"`
}
