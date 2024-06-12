package model

import (
	"github.com/uptrace/bun"
	"time"
)

type InvMastLot struct {
	bun.BaseModel             `bun:"table:inv_mast_lot"`
	InvMastLotUid             int32     `bun:"inv_mast_lot_uid,type:int,pk"`
	InvMastUid                int32     `bun:"inv_mast_uid,type:int"`
	AutoAssignLots            string    `bun:"auto_assign_lots,type:char,default:('N')"`
	AssignmentOption          int32     `bun:"assignment_option,type:int,default:(938)"`
	LotAssignmentRequired     string    `bun:"lot_assignment_required,type:char,default:('N')"`
	DateCreated               time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	DateLastModified          time.Time `bun:"date_last_modified,type:datetime,nullzero"`
	LastMaintainedBy          string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	UseLotCostAsInventoryCost string    `bun:"use_lot_cost_as_inventory_cost,type:char,default:('N')"`
	UseSystemSettingLotAssign string    `bun:"use_system_setting_lot_assign,type:char"`
	LotAttributeGroupUid      int32     `bun:"lot_attribute_group_uid,type:int,nullzero"`
	ShelfLife                 int32     `bun:"shelf_life,type:int,nullzero"`
	ShelfLifeUnitCd           int32     `bun:"shelf_life_unit_cd,type:int,nullzero"`
	ShelfLifeGuar             int32     `bun:"shelf_life_guar,type:int,nullzero"`
	ShelfLifeGuarUnitCd       int32     `bun:"shelf_life_guar_unit_cd,type:int,nullzero"`
	ThirtyDayMonthFlag        string    `bun:"thirty_day_month_flag,type:char,nullzero"`
	ShelfLifeOriginCd         int32     `bun:"shelf_life_origin_cd,type:int,nullzero"`
	ExpDateCalcMethodCd       int32     `bun:"exp_date_calc_method_cd,type:int,nullzero"`
	TrackLotShelfLifeFlag     string    `bun:"track_lot_shelf_life_flag,type:char,nullzero"`
	ShelfLifeGuarPct          float64   `bun:"shelf_life_guar_pct,type:decimal(19,9),nullzero"`
	UseLastCustomerLotFlag    string    `bun:"use_last_customer_lot_flag,type:char,default:('N')"`
	ItemLotExpWarningDays     float64   `bun:"item_lot_exp_warning_days,type:decimal(19,0),default:((0))"`
	BeltItemFlag              string    `bun:"belt_item_flag,type:char,default:('N')"`
	BeltRemnantLength         float64   `bun:"belt_remnant_length,type:decimal(19,9),default:((0))"`
	BeltRemnantWidth          float64   `bun:"belt_remnant_width,type:decimal(19,9),default:((0))"`
	MinimumBeltWidthIncrement float64   `bun:"minimum_belt_width_increment,type:decimal(19,9),nullzero"`
	PurchasingBeltWidth       float64   `bun:"purchasing_belt_width,type:decimal(19,9),nullzero"`
	PurchasingBeltLength      float64   `bun:"purchasing_belt_length,type:decimal(19,9),nullzero"`
	UseLotPricingFlag         string    `bun:"use_lot_pricing_flag,type:char,nullzero"`
	ScrapLimit                float64   `bun:"scrap_limit,type:decimal(19,9),nullzero"`
}
