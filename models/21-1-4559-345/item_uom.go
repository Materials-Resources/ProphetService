package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type ItemUom struct {
	bun.BaseModel    `bun:"table:item_uom"`
	UnitOfMeasure    string    `bun:"unit_of_measure,type:varchar(8),unique"`                        // What is the unit of measure for this row?
	DeleteFlag       string    `bun:"delete_flag,type:char(1)"`                                      // Indicates whether this record is logically deleted
	DateCreated      time.Time `bun:"date_created,type:datetime"`                                    // Indicates the date/time this record was created.
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`                              // Indicates the date/time this record was last modified.
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"` // ID of the user who last maintained this record
	UnitSize         float64   `bun:"unit_size,type:decimal(19,4)"`                                  // Quantity of SKUs in this UOM.
	SellingUnit      *string   `bun:"selling_unit,type:char(1)"`                                     // What is the normal unit of measure that this item
	PurchasingUnit   *string   `bun:"purchasing_unit,type:char(1)"`                                  // What unit of measure is this item purchased in?
	InvMastUid       int32     `bun:"inv_mast_uid,type:int,unique"`                                  // Unique identifier for the item id.
	CreatedBy        *string   `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	ItemUomUid       int32     `bun:"item_uom_uid,type:int,pk"`
	B2bUnitFlag      string    `bun:"b2b_unit_flag,type:char(1),default:('Y')"` // Custom: determines if this UOM may be used on the user's B2B website.
	TallyFactor      *float64  `bun:"tally_factor,type:decimal(19,9)"`          // Custom, factor used to calculate Square Feet Order Quantity from the Tally pop up window.
	WwmsFlag         *string   `bun:"wwms_flag,type:char(1),default:('N')"`     // WWMS unit of measure indicator
	ProdOrderFactor  *int32    `bun:"prod_order_factor,type:int"`               // Custom, factor used for prod order unit conversion.
	MinimumOrderQty  *float64  `bun:"minimum_order_qty,type:decimal(19,4)"`     // Minimum order quantity in terms of the base unit
}
