package model

import (
	"github.com/uptrace/bun"
	"time"
)

type ItemUom struct {
	bun.BaseModel    `bun:"table:item_uom"`
	UnitOfMeasure    string    `bun:"unit_of_measure,type:varchar(8)"`
	DeleteFlag       string    `bun:"delete_flag,type:char"`
	DateCreated      time.Time `bun:"date_created,type:datetime"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"`
	UnitSize         float64   `bun:"unit_size,type:decimal(19,4)"`
	SellingUnit      string    `bun:"selling_unit,type:char,nullzero"`
	PurchasingUnit   string    `bun:"purchasing_unit,type:char,nullzero"`
	InvMastUid       int32     `bun:"inv_mast_uid,type:int"`
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	ItemUomUid       int32     `bun:"item_uom_uid,type:int,pk"`
	B2bUnitFlag      string    `bun:"b2b_unit_flag,type:char,default:('Y')"`
	TallyFactor      float64   `bun:"tally_factor,type:decimal(19,9),nullzero"`
	WwmsFlag         string    `bun:"wwms_flag,type:char,default:('N')"`
	ProdOrderFactor  int32     `bun:"prod_order_factor,type:int,nullzero"`
	MinimumOrderQty  float64   `bun:"minimum_order_qty,type:decimal(19,4),nullzero"`
}
