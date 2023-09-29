package model

import (
	"database/sql"

	"github.com/uptrace/bun"
)

type InvMast struct {
	bun.BaseModel `bun:"table:inv_mast"`
	InvMastUid    int32           `bun:"inv_mast_uid,type:int,pk"`
	ItemId        string          `bun:"item_id,type:varchar(40),unique"`
	ItemDesc      string          `bun:"item_desc,type:varchar(40)"`
	ExtendedDesc  sql.NullString  `bun:"extended_desc,type:varchar(255)"`
	Price1        sql.NullFloat64 `bun:"price1,type:decimal(19, 9)"`
}
