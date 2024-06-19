package gen

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
	DefaultToOrder           string    `bun:"default_to_order,type:char(1),nullzero"`
	RoundingFactor           float64   `bun:"rounding_factor,type:decimal(6,4),nullzero"`
	RequirementType          string    `bun:"requirement_type,type:char(1),nullzero"`
	IncludeUnapprovedPo      string    `bun:"include_unapproved_po,type:char(1),nullzero"`
	IncludeReleaseSchedule   string    `bun:"include_release_schedule,type:char(1),nullzero"`
	ExcludeTboQuantity       string    `bun:"exclude_tbo_quantity,type:char(1),nullzero"`
	ItemUid                  int32     `bun:"item_uid,type:int,nullzero"`
	Roai                     float64   `bun:"roai,type:decimal(19,9),nullzero"`
	NumberOfPeriods          float64   `bun:"number_of_periods,type:decimal(4,0),nullzero"`
	DiscountPercentage       float64   `bun:"discount_percentage,type:decimal(7,4),nullzero"`
	CarryingCost             float64   `bun:"carrying_cost,type:decimal(19,9),nullzero"`
	GporFlag                 string    `bun:"gpor_flag,type:char(1),nullzero"`
	ReturnAllItems           string    `bun:"return_all_items,type:char(1),nullzero"`
	UseSurplusQty            string    `bun:"use_surplus_qty,type:char(1),nullzero"`
	IncludeNonStock          string    `bun:"include_non_stock,type:char(1),nullzero"`
	HighVelocityLevel        float64   `bun:"high_velocity_level,type:decimal(19,4),nullzero"`
	MidVelocityLevel         float64   `bun:"mid_velocity_level,type:decimal(19,4),nullzero"`
	SafetyStockDays          int32     `bun:"safety_stock_days,type:int,nullzero"`
	UsageFactor              float64   `bun:"usage_factor,type:decimal(19,4),nullzero"`
	AutoGenerateTbo          string    `bun:"auto_generate_tbo,type:char(1),nullzero"`
	SummaryDetailInd         string    `bun:"summary_detail_ind,type:char(1),nullzero"`
	BuyerId                  int32     `bun:"buyer_id,type:int,nullzero"`
	UseSurplusFlag           string    `bun:"use_surplus_flag,type:char(1),nullzero"`
	ReplenishFromReplLocOnly string    `bun:"replenish_from_repl_loc_only,type:char(1),nullzero"`
	AutoCalcExpediteDate     string    `bun:"auto_calc_expedite_date,type:char(1),nullzero"`
	DetailCount              int32     `bun:"detail_count,type:int"`
	DateCreated              time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                string    `bun:"created_by,type:varchar(40),default:(suser_sname())"`
	DateLastModified         time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy         string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	UseNsAsSourceForTransfer string    `bun:"use_ns_as_source_for_transfer,type:char(1),nullzero"`
	CombineSpWithStockFlag   string    `bun:"combine_sp_with_stock_flag,type:char(1),default:('N')"`  // Flags to use combine specials and stock items on same PO functionality
	SafetyStockType          int32     `bun:"safety_stock_type,type:int,nullzero"`                    // Type of safety stock in days or service level
	ServiceLevelMeasure      int32     `bun:"service_level_measure,type:int,nullzero"`                // Service measure stock out or backorder
	ServiceLevelPctGoal      float64   `bun:"service_level_pct_goal,type:decimal(19,2),nullzero"`     // Percent of goal for service level
	SupplierIdList           string    `bun:"supplier_id_list,type:text,nullzero"`                    // List of distinct supplier IDs used.
	ProductGroupIdList       string    `bun:"product_group_id_list,type:text,nullzero"`               // List of distinct product group IDs used.
	MinDynamicLookAheadDate  time.Time `bun:"min_dynamic_look_ahead_date,type:datetime,nullzero"`     // The minimum look ahead date to use when using the dynamic look ahead date option
	ParentGporRunHdrUid      int32     `bun:"parent_gpor_run_hdr_uid,type:int,nullzero"`              // Links a gpor_run_hdr record to a an original parent PORG run. Currently used only for custom functionality.
	DeviationLookbackPds     int32     `bun:"deviation_lookback_pds,type:int,nullzero"`               // When calculating safety stock by Deviation Multiplier, the number of periods to look backwards to determine the deviation.  The gpor_run_hdr value tracks what was specified on the Factors tab when requirements were run.
	DeviationMultOnePd       float64   `bun:"deviation_mult_one_pd,type:decimal(19,9),nullzero"`      // When calculating safety stock by Deviation Multiplier, the deviation multiplier to use when one of the lookback periods has forecast usage less than actual usage.  The gpor_run_hdr value tracks what was specified on the Factors tab when requirements were
	DeviationMultTwoPd       float64   `bun:"deviation_mult_two_pd,type:decimal(19,9),nullzero"`      // When calculating safety stock by Deviation Multiplier, the deviation multiplier to use when two of the lookback periods has forecast usage less than actual usage.  The gpor_run_hdr value tracks what was specified on the Factors tab when requirements were
	DeviationMultThreePd     float64   `bun:"deviation_mult_three_pd,type:decimal(19,9),nullzero"`    // When calculating safety stock by Deviation Multiplier, the deviation multiplier to use when three or more of the lookback periods has forecast usage less than actual usage.  The gpor_run_hdr value tracks what was specified on the Factors tab when requirem
	MinSafetyStockDays       float64   `bun:"min_safety_stock_days,type:decimal(19,9),nullzero"`      // Minimum fence of safety stock days to use which overrides safety stock days calculation.  The gpor_run_hdr value tracks what was specified on the Factors tab when requirements were run.
	MaxSafetyStockDays       float64   `bun:"max_safety_stock_days,type:decimal(19,9),nullzero"`      // Maximum fence of safety stock days to use which overrides safety stock days calculation.  The gpor_run_hdr value tracks what was specified on the Factors tab when requirements were run.
	CritItemDeviationMult    float64   `bun:"crit_item_deviation_mult,type:decimal(19,9),nullzero"`   // When calculating safety stock by Deviation Multiplier, the deviation multiplier to use when the item/location is flagged as a critical item/location.  The gpor_run_hdr value tracks what was specified on the Factors tab when requirements were run.
	CritMinSafetyStockDays   float64   `bun:"crit_min_safety_stock_days,type:decimal(19,9),nullzero"` // Minimum fence of safety stock days to use which overrides safety stock days calculation when item/location is marked as a critical item/location.  The gpor_run_hdr value tracks what was specified on the Factors tab when requirements were run.
	CritMaxSafetyStockDays   float64   `bun:"crit_max_safety_stock_days,type:decimal(19,9),nullzero"` // Maximum fence of safety stock days to use which overrides safety stock days calculation when item/location is marked as a critical item/location.  The gpor_run_hdr value tracks what was specified on the Factors tab when requirements were run.
}
