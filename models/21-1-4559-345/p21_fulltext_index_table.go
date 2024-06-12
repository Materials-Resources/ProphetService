package model

import "github.com/uptrace/bun"

type P21FulltextIndexTable struct {
	bun.BaseModel            `bun:"table:p21_fulltext_index_table"`
	P21FulltextIndexTableUid int32  `bun:"p21_fulltext_index_table_uid,type:int,pk,identity"`
	P21FulltextCatalogUid    int32  `bun:"p21_fulltext_catalog_uid,type:int"`
	TableName                string `bun:"table_name,type:varchar(256)"`
	PrimaryKey               string `bun:"primary_key,type:varchar(256)"`
}
