package gen

import "github.com/uptrace/bun"

type P21FulltextIndexTable struct {
	bun.BaseModel            `bun:"table:p21_fulltext_index_table"`
	P21FulltextIndexTableUid int32  `bun:"p21_fulltext_index_table_uid,type:int,autoincrement,pk"` // Unique identifier for p21_fulltext_index_table
	P21FulltextCatalogUid    int32  `bun:"p21_fulltext_catalog_uid,type:int"`                      // The Fulltext Catalog that this table belongs to.
	TableName                string `bun:"table_name,type:varchar(256)"`                           // Name of the table that will have a Fulltext Index
	PrimaryKey               string `bun:"primary_key,type:varchar(256)"`                          // Each Fulltext Index needs a primary key. This holds the primary key that the index will use for this table.
}
