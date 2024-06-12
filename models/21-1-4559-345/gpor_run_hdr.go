package model

import (
	"github.com/uptrace/bun"
	"time"
)

type GporRunHdr struct {
	bun.BaseModel            `bun:"table:gpor_run_hdr"`
	GporRunHdrUid            int32     `bun:"gpor_run_hdr_uid,type:int,pk"`
	PurchasingLocationId     int32     `bun:"purchasing_location_id,type:int,nullzero"`
	PurchaseTransferGroupId  string    `bun:"Purchase_transfer_group_id,type:char(20),nullzero"`
	BegSupplierId            int32     `bun:"beg_supplier_id,type:int,nullzero"`
	EndSupplierId            int32     `bun:"end_supplier_id,type:int,nullzero"`
	BegAbcClassId            string    `bun:"beg_ABC_class_id,type:char(40),nullzero"`
	EndAbcClassId            string    `bun:"end_abc_class_id,type:char(40),nullzero"`
	BegProductGroupId        string    `bun:"beg_product_group_id,type:char(40),nullzero"`
	EndProductGroupId        string    `bun:"end_product_group_id,type:char(40),nullzero"`
	LookAheadDate            time.Time `bun:"look_ahead_date,type:datetime,nullzero"`
	ReplenishmentMethod      string    `bun:"replenishment_method,type:char(10),nullzero"`
	OrderPointException      float64   `bun:"order_point_exception,type:decimal(19,9),nullzero"`
	DefaultToOrder           string    `bun:"default_to_order,type:char,nullzero"`
	RoundingFactor           float64   `bun:"rounding_factor,type:decimal(6,4),nullzero"`
	RequirementType          string    `bun:"requirement_type,type:char,nullzero"`
	IncludeUnapprovedPo      string    `bun:"include_unapproved_po,type:char,nullzero"`
	IncludeReleaseSchedule   string    `bun:"include_release_schedule,type:char,nullzero"`
	ExcludeTboQuantity       string    `bun:"exclude_tbo_quantity,type:char,nullzero"`
	ItemUid                  int32     `bun:"item_uid,type:int,nullzero"`
	Roai                     float64   `bun:"roai,type:decimal(19,9),nullzero"`
	NumberOfPeriods          float64   `bun:"number_of_periods,type:decimal(4,0),nullzero"`
	DiscountPercentage       float64   `bun:"discount_percentage,type:decimal(7,4),nullzero"`
	CarryingCost             float64   `bun:"carrying_cost,type:decimal(19,9),nullzero"`
	GporFlag                 string    `bun:"gpor_flag,type:char,nullzero"`
	ReturnAllItems           string    `bun:"return_all_items,type:char,nullzero"`
	UseSurplusQty            string    `bun:"use_surplus_qty,type:char,nullzero"`
	IncludeNonStock          string    `bun:"include_non_stock,type:char,nullzero"`
	HighVelocityLevel        float64   `bun:"high_velocity_level,type:decimal(19,4),nullzero"`
	MidVelocityLevel         float64   `bun:"mid_velocity_level,type:decimal(19,4),nullzero"`
	SafetyStockDays          int32     `bun:"safety_stock_days,type:int,nullzero"`
	UsageFactor              float64   `bun:"usage_factor,type:decimal(19,4),nullzero"`
	AutoGenerateTbo          string    `bun:"auto_generate_tbo,type:char,nullzero"`
	SummaryDetailInd         string    `bun:"summary_detail_ind,type:char,nullzero"`
	BuyerId                  int32     `bun:"buyer_id,type:int,nullzero"`
	UseSurplusFlag           string    `bun:"use_surplus_flag,type:char,nullzero"`
	ReplenishFromReplLocOnly string    `bun:"replenish_from_repl_loc_only,type:char,nullzero"`
	AutoCalcExpediteDate     string    `bun:"auto_calc_expedite_date,type:char,nullzero"`
	DetailCount              int32     `bun:"detail_count,type:int"`
	DateCreated              time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                string    `bun:"created_by,type:varchar(40),default:(suser_sname())"`
	DateLastModified         time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy         string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	UseNsAsSourceForTransfer string    `bun:"use_ns_as_source_for_transfer,type:char,nullzero"`
	CombineSpWithStockFlag   string    `bun:"combine_sp_with_stock_flag,type:char,default:('N')"`
	SafetyStockType          int32     `bun:"safety_stock_type,type:int,nullzero"`
	ServiceLevelMeasure      int32     `bun:"service_level_measure,type:int,nullzero"`
	ServiceLevelPctGoal      float64   `bun:"service_level_pct_goal,type:decimal(19,2),nullzero"`
	SupplierIdList           string    `bun:"supplier_id_list,type:text(2147483647),nullzero"`
	ProductGroupIdList       string    `bun:"product_group_id_list,type:text(2147483647),nullzero"`
	MinDynamicLookAheadDate  time.Time `bun:"min_dynamic_look_ahead_date,type:datetime,nullzero"`
	ParentGporRunHdrUid      int32     `bun:"parent_gpor_run_hdr_uid,type:int,nullzero"`
	DeviationLookbackPds     int32     `bun:"deviation_lookback_pds,type:int,nullzero"`
	DeviationMultOnePd       float64   `bun:"deviation_mult_one_pd,type:decimal(19,9),nullzero"`
	DeviationMultTwoPd       float64   `bun:"deviation_mult_two_pd,type:decimal(19,9),nullzero"`
	DeviationMultThreePd     float64   `bun:"deviation_mult_three_pd,type:decimal(19,9),nullzero"`
	MinSafetyStockDays       float64   `bun:"min_safety_stock_days,type:decimal(19,9),nullzero"`
	MaxSafetyStockDays       float64   `bun:"max_safety_stock_days,type:decimal(19,9),nullzero"`
	CritItemDeviationMult    float64   `bun:"crit_item_deviation_mult,type:decimal(19,9),nullzero"`
	CritMinSafetyStockDays   float64   `bun:"crit_min_safety_stock_days,type:decimal(19,9),nullzero"`
	CritMaxSafetyStockDays   float64   `bun:"crit_max_safety_stock_days,type:decimal(19,9),nullzero"`
}
