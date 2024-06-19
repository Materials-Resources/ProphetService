package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type PurchaseClass struct {
	bun.BaseModel               `bun:"table:purchase_class"`
	PurchaseClassId             string    `bun:"purchase_class_id,type:varchar(8),pk"`                      // What is the unique identifier for this purchase cl
	PurchaseClassDescription    string    `bun:"purchase_class_description,type:varchar(40)"`               // What is this purchase class for?
	DeleteFlag                  string    `bun:"delete_flag,type:char(1)"`                                  // Indicates whether this record is logically deleted
	DateCreated                 time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified            time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy            string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	HighLevelMonths             float64   `bun:"high_level_months,type:decimal(19,2),nullzero"`             // Order up to levels in terms of month’s supply for high velocity items.
	MidLevelMonths              float64   `bun:"mid_level_months,type:decimal(19,2),nullzero"`              // Order up to levels in terms of month’s supply for medium velocity items.
	LowLevelMonths              float64   `bun:"low_level_months,type:decimal(19,2),nullzero"`              // Order up to levels in terms of month’s supply for low velocity items.
	AnnualThreshold             float64   `bun:"annual_threshold,type:decimal(19,4),nullzero"`              // Annual cost used for ranking purposes.
	HighSafetyStockFactor       float64   `bun:"high_safety_stock_factor,type:decimal(19,4),nullzero"`      // Safety stock factor (applied against the safety days) for high volume items.
	MidSafetyStockFactor        float64   `bun:"mid_safety_stock_factor,type:decimal(19,4),default:('N')"`  // Safety stock factor (applied against the safety days) for medium volume items.
	LowSafetyStockFactor        float64   `bun:"low_safety_stock_factor,type:decimal(19,4),default:('N')"`  // Safety stock factor (applied against the safety days) for low volume items.
	ExcludeFromRanking          string    `bun:"exclude_from_ranking,type:char(1)"`                         // Is this purchase class excluded from inventory ranking?
	ExcludeFromReclassification string    `bun:"exclude_from_reclassification,type:char(1)"`                // Is this purchase class excluded from inventory reclassification?
	SplittableFlag              string    `bun:"splittable_flag,type:char(1),default:('N')"`                // Determines whether or not this class allows OE line splitting.
	SafetyStockType             int32     `bun:"safety_stock_type,type:int,nullzero"`                       // Method used to determine safety stock
	ServiceLevelMeasure         int32     `bun:"service_level_measure,type:int,nullzero"`                   // Customer service level (stock out back order)
	ServiceLevelPctGoal         float64   `bun:"service_level_pct_goal,type:decimal(19,2),nullzero"`        // Percent of service level goal for safety stock
	SafetyStockDays             float64   `bun:"safety_stock_days,type:decimal(19,4),default:((0.0000))"`   // Number of days worth of safety stock
	ErraticMaxFactor            float64   `bun:"erratic_max_factor,type:decimal(19,4),nullzero"`            // Used in min/max identification.  Multiplied against the new min for erratic items to determine the new max.
	NonerraticMinFactor         float64   `bun:"nonerratic_min_factor,type:decimal(19,4),nullzero"`         // Used in min/max identification.  Multiplied against the purchasing unit size for nonerratic items to suggest a new min.
	NonerraticMaxFactor         float64   `bun:"nonerratic_max_factor,type:decimal(19,4),nullzero"`         // Used in min/max identification.  Multiplied against the new min for nonerratic items to determine the new max.
	IncludeLargeOrderQtyFlag    string    `bun:"include_large_order_qty_flag,type:char(1),nullzero"`        // Custom F42518: determines if item is assigned this purchase class are included in the OE unusually large order qty functionality
	MinimumPercentUsed          float64   `bun:"minimum_percent_used,type:decimal(19,4),default:((0))"`     // If forecast is less than this % of actual usage, then set filtered usage to forecast
	MaximumPercentUsed          float64   `bun:"maximum_percent_used,type:decimal(19,4),default:((0))"`     // If forecast is more than this % of actual usage, then set filtered usage to forecast
	DeviationLookbackPds        int32     `bun:"deviation_lookback_pds,type:int,nullzero"`                  // When calculating safety stock by Deviation Multiplier, the number of periods to look backwards to determine the deviation.
	DeviationMultOnePd          float64   `bun:"deviation_mult_one_pd,type:decimal(19,9),nullzero"`         // When calculating safety stock by Deviation Multiplier, the deviation multiplier to use when one of the lookback periods has forecast usage less than actual usage.
	DeviationMultTwoPd          float64   `bun:"deviation_mult_two_pd,type:decimal(19,9),nullzero"`         // When calculating safety stock by Deviation Multiplier, the deviation multiplier to use when two of the lookback periods has forecast usage less than actual usage.
	DeviationMultThreePd        float64   `bun:"deviation_mult_three_pd,type:decimal(19,9),nullzero"`       // When calculating safety stock by Deviation Multiplier, the deviation multiplier to use when three or more of the lookback periods has forecast usage less than actual usage.
	MinSafetyStockDays          float64   `bun:"min_safety_stock_days,type:decimal(19,9),nullzero"`         // Minimum fence of safety stock days to use which overrides safety stock days calculation.
	MaxSafetyStockDays          float64   `bun:"max_safety_stock_days,type:decimal(19,9),nullzero"`         // Maximum fence of safety stock days to use which overrides safety stock days calculation.
	CritItemDeviationMult       float64   `bun:"crit_item_deviation_mult,type:decimal(19,9),nullzero"`      // When calculating safety stock by Deviation Multiplier, the deviation multiplier to use when the item/location is flagged as a critical item/location.
	CritMinSafetyStockDays      float64   `bun:"crit_min_safety_stock_days,type:decimal(19,9),nullzero"`    // Minimum fence of safety stock days to use which overrides safety stock days calculation when item/location is marked as a critical item/location.
	CritMaxSafetyStockDays      float64   `bun:"crit_max_safety_stock_days,type:decimal(19,9),nullzero"`    // Maximum fence of safety stock days to use which overrides safety stock days calculation when item/location is marked as a critical item/location.
}
