package model

import (
	"database/sql"

	"github.com/uptrace/bun"
)

type OeLine struct {
	bun.BaseModel `bun:"table:oe_line"`
	OrderNo       string          `bun:"order_no,type:varchar(8),pk"`
	LineNo        float64         `bun:"line_no,type:decimal(19),pk"`
	UnitPrice     sql.NullFloat64 `bun:"unit_price,type:decimal(19, 9)"`
	QtyOrdered    sql.NullFloat64 `bun:"qty_ordered,type:decimal(19, 9)"`
	ExtendedPrice sql.NullFloat64 `bun:"extended_price,type:decimal(19, 4)"`
	UnitOfMeasure sql.NullString  `bun:"unit_of_measure,type:varchar(8)"`
	InvMastUid    int32           `bun:"inv_mast_uid,type:int"`

	InvMast InvMast `bun:"rel:has-one,join:inv_mast_uid=inv_mast_uid"`
}
