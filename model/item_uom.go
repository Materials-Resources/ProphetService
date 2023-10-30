package model

import (
	"database/sql"
	"time"

	"github.com/uptrace/bun"
)

type ItemUom struct {
	bun.BaseModel `bun:"table:item_uom"`

	UnitOfMeasure    string          `bun:"unit_of_measure,type:varchar(8)"`
	DeleteFlag       string          `bun:"delete_flag,type:char"`
	DateCreated      time.Time       `bun:"date_created,type:datetime"`
	DateLastModified time.Time       `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy string          `bun:"last_maintained_by,type:varchar(30)"`
	UnitSize         float64         `bun:"unit_size,type:decimal(19, 4)"`
	SellingUnit      sql.NullString  `bun:"selling_unit,type:char"`
	PurchasingUnit   sql.NullString  `bun:"purchasing_unit,type:char"`
	InvMastUid       int64           `bun:"inv_mast_uid,type:int"`
	CreatedBy        sql.NullString  `bun:"created_by,type:varchar(255)"`
	ItemUomUid       int32           `bun:"item_uom_uid,type:int,pk"`
	B2BUnitFlag      string          `bun:"b2b_unit_flag,type:char"`
	TallyFactor      sql.NullFloat64 `bun:"tally_factor,type:decimal(19, 9)"`
	WwmsFlag         string          `bun:"wwms_flag,type:char"`
	ProdOrderFactor  sql.NullFloat64 `bun:"prod_order_factor,type:char"`
	MinimumOrderQty  sql.NullFloat64 `bun:"minimum_order_qty,type:int"`
}
