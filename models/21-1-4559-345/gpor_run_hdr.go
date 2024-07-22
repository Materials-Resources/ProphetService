package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type GporRunHdr struct {
	bun.BaseModel            `bun:"table:gpor_run_hdr"`
	GporRunHdrUid            int32      `bun:"gpor_run_hdr_uid,type:int,pk"`
	PurchasingLocationId     *int32     `bun:"purchasing_location_id,type:int"`
	PurchaseTransferGroupId  *string    `bun:"Purchase_transfer_group_id,type:char(20)"`
	BegSupplierId            *int32     `bun:"beg_supplier_id,type:int"`
	EndSupplierId            *int32     `bun:"end_supplier_id,type:int"`
	BegAbcClassId            *string    `bun:"beg_ABC_class_id,type:char(40)"`
	EndAbcClassId            *string    `bun:"end_abc_class_id,type:char(40)"`
	BegProductGroupId        *string    `bun:"beg_product_group_id,type:char(40)"`
	EndProductGroupId        *string    `bun:"end_product_group_id,type:char(40)"`
	LookAheadDate            *time.Time `bun:"look_ahead_date,type:datetime"`
	ReplenishmentMethod      *string    `bun:"replenishment_method,type:char(10)"`
	OrderPointException      *float64   `bun:"order_point_exception,type:decimal(19,9)"`
	DefaultToOrder           *string    `bun:"default_to_order,type:char(1)"`
	RoundingFactor           *float64   `bun:"rounding_factor,type:decimal(6,4)"`
	RequirementType          *string    `bun:"requirement_type,type:char(1)"`
	IncludeUnapprovedPo      *string    `bun:"include_unapproved_po,type:char(1)"`
	IncludeReleaseSchedule   *string    `bun:"include_release_schedule,type:char(1)"`
	ExcludeTboQuantity       *string    `bun:"exclude_tbo_quantity,type:char(1)"`
	ItemUid                  *int32     `bun:"item_uid,type:int"`
	Roai                     *float64   `bun:"roai,type:decimal(19,9)"`
	NumberOfPeriods          *float64   `bun:"number_of_periods,type:decimal(4,0)"`
	DiscountPercentage       *float64   `bun:"discount_percentage,type:decimal(7,4)"`
	CarryingCost             *float64   `bun:"carrying_cost,type:decimal(19,9)"`
	GporFlag                 *string    `bun:"gpor_flag,type:char(1)"`
	ReturnAllItems           *string    `bun:"return_all_items,type:char(1)"`
	UseSurplusQty            *string    `bun:"use_surplus_qty,type:char(1)"`
	IncludeNonStock          *string    `bun:"include_non_stock,type:char(1)"`
	HighVelocityLevel        *float64   `bun:"high_velocity_level,type:decimal(19,4)"`
	MidVelocityLevel         *float64   `bun:"mid_velocity_level,type:decimal(19,4)"`
	SafetyStockDays          *int32     `bun:"safety_stock_days,type:int"`
	UsageFactor              *float64   `bun:"usage_factor,type:decimal(19,4)"`
	AutoGenerateTbo          *string    `bun:"auto_generate_tbo,type:char(1)"`
	SummaryDetailInd         *string    `bun:"summary_detail_ind,type:char(1)"`
	BuyerId                  *int32     `bun:"buyer_id,type:int"`
	UseSurplusFlag           *string    `bun:"use_surplus_flag,type:char(1)"`
	ReplenishFromReplLocOnly *string    `bun:"replenish_from_repl_loc_only,type:char(1)"`
	AutoCalcExpediteDate     *string    `bun:"auto_calc_expedite_date,type:char(1)"`
	DetailCount              int32      `bun:"detail_count,type:int"`
	DateCreated              time.Time  `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                string     `bun:"created_by,type:varchar(40),default:(suser_sname())"`
	DateLastModified         time.Time  `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy         string     `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	UseNsAsSourceForTransfer *string    `bun:"use_ns_as_source_for_transfer,type:char(1)"`
	CombineSpWithStockFlag   string     `bun:"combine_sp_with_stock_flag,type:char(1),default:('N')"` // Flags to use combine specials and stock items on same PO functionality
	SafetyStockType          *int32     `bun:"safety_stock_type,type:int"`                            // Type of safety stock in days or service level
	ServiceLevelMeasure      *int32     `bun:"service_level_measure,type:int"`                        // Service measure stock out or backorder
	ServiceLevelPctGoal      *float64   `bun:"service_level_pct_goal,type:decimal(19,2)"`             // Percent of goal for service level
	SupplierIdList           *string    `bun:"supplier_id_list,type:text"`                            // List of distinct supplier IDs used.
	ProductGroupIdList       *string    `bun:"product_group_id_list,type:text"`                       // List of distinct product group IDs used.
	MinDynamicLookAheadDate  *time.Time `bun:"min_dynamic_look_ahead_date,type:datetime"`             // The minimum look ahead date to use when using the dynamic look ahead date option
	ParentGporRunHdrUid      *int32     `bun:"parent_gpor_run_hdr_uid,type:int"`                      // Links a gpor_run_hdr record to a an original parent PORG run. Currently used only for custom functionality.
	DeviationLookbackPds     *int32     `bun:"deviation_lookback_pds,type:int"`                       // When calculating safety stock by Deviation Multiplier, the number of periods to look backwards to determine the deviation.  The gpor_run_hdr value tracks what was specified on the Factors tab when requirements were run.
	DeviationMultOnePd       *float64   `bun:"deviation_mult_one_pd,type:decimal(19,9)"`              // When calculating safety stock by Deviation Multiplier, the deviation multiplier to use when one of the lookback periods has forecast usage less than actual usage.  The gpor_run_hdr value tracks what was specified on the Factors tab when requirements were
	DeviationMultTwoPd       *float64   `bun:"deviation_mult_two_pd,type:decimal(19,9)"`              // When calculating safety stock by Deviation Multiplier, the deviation multiplier to use when two of the lookback periods has forecast usage less than actual usage.  The gpor_run_hdr value tracks what was specified on the Factors tab when requirements were
	DeviationMultThreePd     *float64   `bun:"deviation_mult_three_pd,type:decimal(19,9)"`            // When calculating safety stock by Deviation Multiplier, the deviation multiplier to use when three or more of the lookback periods has forecast usage less than actual usage.  The gpor_run_hdr value tracks what was specified on the Factors tab when requirem
	MinSafetyStockDays       *float64   `bun:"min_safety_stock_days,type:decimal(19,9)"`              // Minimum fence of safety stock days to use which overrides safety stock days calculation.  The gpor_run_hdr value tracks what was specified on the Factors tab when requirements were run.
	MaxSafetyStockDays       *float64   `bun:"max_safety_stock_days,type:decimal(19,9)"`              // Maximum fence of safety stock days to use which overrides safety stock days calculation.  The gpor_run_hdr value tracks what was specified on the Factors tab when requirements were run.
	CritItemDeviationMult    *float64   `bun:"crit_item_deviation_mult,type:decimal(19,9)"`           // When calculating safety stock by Deviation Multiplier, the deviation multiplier to use when the item/location is flagged as a critical item/location.  The gpor_run_hdr value tracks what was specified on the Factors tab when requirements were run.
	CritMinSafetyStockDays   *float64   `bun:"crit_min_safety_stock_days,type:decimal(19,9)"`         // Minimum fence of safety stock days to use which overrides safety stock days calculation when item/location is marked as a critical item/location.  The gpor_run_hdr value tracks what was specified on the Factors tab when requirements were run.
	CritMaxSafetyStockDays   *float64   `bun:"crit_max_safety_stock_days,type:decimal(19,9)"`         // Maximum fence of safety stock days to use which overrides safety stock days calculation when item/location is marked as a critical item/location.  The gpor_run_hdr value tracks what was specified on the Factors tab when requirements were run.
}
