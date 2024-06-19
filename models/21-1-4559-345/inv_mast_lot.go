package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type InvMastLot struct {
	bun.BaseModel             `bun:"table:inv_mast_lot"`
	InvMastLotUid             int32     `bun:"inv_mast_lot_uid,type:int,pk"`                               // Unique identifier for an inv_mast_lot record.
	InvMastUid                int32     `bun:"inv_mast_uid,type:int,unique"`                               // Unique identifier for the item id.
	AutoAssignLots            string    `bun:"auto_assign_lots,type:char(1),default:('N')"`                // Assign lots for allocated material automatically
	AssignmentOption          int32     `bun:"assignment_option,type:int,default:(938)"`                   // Method to auto-assign lots based on auto_assign_lot option (FIFO, lot complete, etc)
	LotAssignmentRequired     string    `bun:"lot_assignment_required,type:char(1),default:('N')"`         // Require lot assignment for allocated quantities
	DateCreated               time.Time `bun:"date_created,type:datetime,default:(getdate())"`             // Indicates the date/time this record was created.
	DateLastModified          time.Time `bun:"date_last_modified,type:datetime,nullzero"`                  // Indicates the date/time this record was last modified.
	LastMaintainedBy          string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`  // ID of the user who last maintained this record
	UseLotCostAsInventoryCost string    `bun:"use_lot_cost_as_inventory_cost,type:char(1),default:('N')"`  // Enable the lot costing feature for this item
	UseSystemSettingLotAssign string    `bun:"use_system_setting_lot_assign,type:char(1)"`                 // Use system settings for the lot assignment options
	LotAttributeGroupUid      int32     `bun:"lot_attribute_group_uid,type:int,nullzero"`                  // determines which lot_attribute_group assigned to this lot tracked item
	ShelfLife                 int32     `bun:"shelf_life,type:int,nullzero"`                               // Shelf life duration.  Meaning depends on shelf_life_unit_cd (days or months).
	ShelfLifeUnitCd           int32     `bun:"shelf_life_unit_cd,type:int,nullzero"`                       // Shelf life duration unit.
	ShelfLifeGuar             int32     `bun:"shelf_life_guar,type:int,nullzero"`                          // Shelf life guarantee duration.  Meaning depends on shelf_life_guar_unit_cd (days or months).
	ShelfLifeGuarUnitCd       int32     `bun:"shelf_life_guar_unit_cd,type:int,nullzero"`                  // Shelf life guarantee duration unit.
	ThirtyDayMonthFlag        string    `bun:"thirty_day_month_flag,type:char(1),nullzero"`                // Indicates if shelf life duration will be calculated using 30 day months (if unit is month).
	ShelfLifeOriginCd         int32     `bun:"shelf_life_origin_cd,type:int,nullzero"`                     // Indicates when shelf life begins.
	ExpDateCalcMethodCd       int32     `bun:"exp_date_calc_method_cd,type:int,nullzero"`                  // Expiration date calculation method.
	TrackLotShelfLifeFlag     string    `bun:"track_lot_shelf_life_flag,type:char(1),nullzero"`            // Indicates if lot shelf life should be tracked for this item.
	ShelfLifeGuarPct          float64   `bun:"shelf_life_guar_pct,type:decimal(19,9),nullzero"`            // Shelf life supplier guarantee percent.  Informational only.  Supplier guaranteess that x% of shelf life will be reamaining upon delivery.
	UseLastCustomerLotFlag    string    `bun:"use_last_customer_lot_flag,type:char(1),default:('N')"`      // Use last lot sent to customer when assigning lots if possible
	ItemLotExpWarningDays     float64   `bun:"item_lot_exp_warning_days,type:decimal(19,0),default:((0))"` // Number of days before lot expires to warn user
	BeltItemFlag              string    `bun:"belt_item_flag,type:char(1),default:('N')"`                  // Indicates if this lot item is a belt item
	BeltRemnantLength         float64   `bun:"belt_remnant_length,type:decimal(19,9),default:((0))"`       // The height which when reached the lot will be marked as a remnant
	BeltRemnantWidth          float64   `bun:"belt_remnant_width,type:decimal(19,9),default:((0))"`        // The width which when reached the lot will be marked as a remnant
	MinimumBeltWidthIncrement float64   `bun:"minimum_belt_width_increment,type:decimal(19,9),nullzero"`   // Measurement in inches that the customer would be charged for.
	PurchasingBeltWidth       float64   `bun:"purchasing_belt_width,type:decimal(19,9),nullzero"`          // Measurement that indicates if no belt pieces are this wide,  the belt should be purchased.
	PurchasingBeltLength      float64   `bun:"purchasing_belt_length,type:decimal(19,9),nullzero"`         // Measurement that indicates if no belt pieces are this long,  the belt should be purchased.
	UseLotPricingFlag         string    `bun:"use_lot_pricing_flag,type:char(1),nullzero"`                 // Indicate if item will be priced on a lot by lot basis via the multiplier
	ScrapLimit                float64   `bun:"scrap_limit,type:decimal(19,9),nullzero"`                    // Scrap limit for lot assignment
}
