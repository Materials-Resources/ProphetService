package model

import "github.com/uptrace/bun"

type P21FulltextCatalog struct {
	bun.BaseModel         `bun:"table:p21_fulltext_catalog"`
	P21FulltextCatalogUid int32  `bun:"p21_fulltext_catalog_uid,type:int,pk,identity"`
	CatalogName           string `bun:"catalog_name,type:varchar(256)"`
	Counter               int16  `bun:"counter,type:smallint,default:((0))"`
}
