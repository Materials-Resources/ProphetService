package model

import (
	"github.com/uptrace/bun"
	"time"
)

type PurchaseClass struct {
	bun.BaseModel               `bun:"table:purchase_class"`
	PurchaseClassId             string    `bun:"purchase_class_id,type:varchar(8),pk"`
	PurchaseClassDescription    string    `bun:"purchase_class_description,type:varchar(40)"`
	DeleteFlag                  string    `bun:"delete_flag,type:char"`
	DateCreated                 time.Time `bun:"date_created,type:datetime"`
	DateLastModified            time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy            string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	HighLevelMonths             float64   `bun:"high_level_months,type:decimal(19,2),nullzero"`
	MidLevelMonths              float64   `bun:"mid_level_months,type:decimal(19,2),nullzero"`
	LowLevelMonths              float64   `bun:"low_level_months,type:decimal(19,2),nullzero"`
	AnnualThreshold             float64   `bun:"annual_threshold,type:decimal(19,4),nullzero"`
	HighSafetyStockFactor       float64   `bun:"high_safety_stock_factor,type:decimal(19,4),nullzero"`
	MidSafetyStockFactor        float64   `bun:"mid_safety_stock_factor,type:decimal(19,4),default:('N')"`
	LowSafetyStockFactor        float64   `bun:"low_safety_stock_factor,type:decimal(19,4),default:('N')"`
	ExcludeFromRanking          string    `bun:"exclude_from_ranking,type:char"`
	ExcludeFromReclassification string    `bun:"exclude_from_reclassification,type:char"`
	SplittableFlag              string    `bun:"splittable_flag,type:char,default:('N')"`
	SafetyStockType             int32     `bun:"safety_stock_type,type:int,nullzero"`
	ServiceLevelMeasure         int32     `bun:"service_level_measure,type:int,nullzero"`
	ServiceLevelPctGoal         float64   `bun:"service_level_pct_goal,type:decimal(19,2),nullzero"`
	SafetyStockDays             float64   `bun:"safety_stock_days,type:decimal(19,4),default:((0.0000))"`
	ErraticMaxFactor            float64   `bun:"erratic_max_factor,type:decimal(19,4),nullzero"`
	NonerraticMinFactor         float64   `bun:"nonerratic_min_factor,type:decimal(19,4),nullzero"`
	NonerraticMaxFactor         float64   `bun:"nonerratic_max_factor,type:decimal(19,4),nullzero"`
	IncludeLargeOrderQtyFlag    string    `bun:"include_large_order_qty_flag,type:char,nullzero"`
	MinimumPercentUsed          float64   `bun:"minimum_percent_used,type:decimal(19,4),default:((0))"`
	MaximumPercentUsed          float64   `bun:"maximum_percent_used,type:decimal(19,4),default:((0))"`
	DeviationLookbackPds        int32     `bun:"deviation_lookback_pds,type:int,nullzero"`
	DeviationMultOnePd          float64   `bun:"deviation_mult_one_pd,type:decimal(19,9),nullzero"`
	DeviationMultTwoPd          float64   `bun:"deviation_mult_two_pd,type:decimal(19,9),nullzero"`
	DeviationMultThreePd        float64   `bun:"deviation_mult_three_pd,type:decimal(19,9),nullzero"`
	MinSafetyStockDays          float64   `bun:"min_safety_stock_days,type:decimal(19,9),nullzero"`
	MaxSafetyStockDays          float64   `bun:"max_safety_stock_days,type:decimal(19,9),nullzero"`
	CritItemDeviationMult       float64   `bun:"crit_item_deviation_mult,type:decimal(19,9),nullzero"`
	CritMinSafetyStockDays      float64   `bun:"crit_min_safety_stock_days,type:decimal(19,9),nullzero"`
	CritMaxSafetyStockDays      float64   `bun:"crit_max_safety_stock_days,type:decimal(19,9),nullzero"`
}
