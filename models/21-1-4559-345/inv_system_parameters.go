package model

import (
	"github.com/uptrace/bun"
	"time"
)

type InvSystemParameters struct {
	bun.BaseModel             `bun:"table:inv_system_parameters"`
	UseBinCapability          string    `bun:"use_bin_capability,type:char"`
	DateCreated               time.Time `bun:"date_created,type:datetime"`
	DateLastModified          time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy          string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"`
	NumberOfDecimalPlaces     int32     `bun:"number_of_decimal_places,type:int,nullzero"`
	InvScan1                  float64   `bun:"inv_scan_1,type:decimal(1,0),nullzero"`
	InvScan2                  float64   `bun:"inv_scan_2,type:decimal(1,0),nullzero"`
	InvScan3                  float64   `bun:"inv_scan_3,type:decimal(1,0),nullzero"`
	InvScan4                  float64   `bun:"inv_scan_4,type:decimal(1,0),nullzero"`
	InvScan5                  float64   `bun:"inv_scan_5,type:decimal(1,0),nullzero"`
	InvScan6                  float64   `bun:"inv_scan_6,type:decimal(1,0),nullzero"`
	InventoryCostingBasis     string    `bun:"inventory_costing_basis,type:varchar(16)"`
	HighAvgMonthlyUsage       float64   `bun:"high_avg_monthly_usage,type:decimal(19,4),nullzero"`
	MidAvgMonthlyUsage        float64   `bun:"mid_avg_monthly_usage,type:decimal(19,4),nullzero"`
	LowAvgMonthlyUsage        float64   `bun:"low_avg_monthly_usage,type:decimal(19,4),nullzero"`
	HighVelocityLevel         float64   `bun:"high_velocity_level,type:decimal(19,4),nullzero"`
	MidVelocityLevel          float64   `bun:"mid_velocity_level,type:decimal(19,4),nullzero"`
	NumberOfDays              float64   `bun:"number_of_days,type:decimal(19,0),nullzero"`
	MinExceptionalSale        float64   `bun:"min_exceptional_sale,type:decimal(19,4)"`
	ExceptionalSalesPercent   float64   `bun:"exceptional_sales_percent,type:decimal(19,4)"`
	OrderHitPercent           float64   `bun:"order_hit_percent,type:decimal(19,4)"`
	DefaultUom                string    `bun:"default_uom,type:char,nullzero"`
	MaintainMovingAverageCost string    `bun:"maintain_moving_average_cost,type:char,nullzero"`
	ExcludeDueInFromNetstock  string    `bun:"exclude_due_in_from_netstock,type:char"`
	RoundingFactor            float64   `bun:"rounding_factor,type:decimal(5,4),nullzero"`
	AllowUncostedAdjustments  string    `bun:"allow_uncosted_adjustments,type:char"`
	DefaultAdjustmentCost     string    `bun:"default_adjustment_cost,type:char"`
	ReceiptsForLeadTime       int32     `bun:"receipts_for_lead_time,type:int,default:(4)"`
	InvSystemParameterUid     int32     `bun:"inv_system_parameter_uid,type:int,pk"`
	UseLotCapability          string    `bun:"use_lot_capability,type:char,default:('N')"`
	UseSerialCapability       string    `bun:"use_serial_capability,type:char,default:('N')"`
	TrackSerialDimensionsCd   string    `bun:"track_serial_dimensions_cd,type:char,default:('N')"`
	DimensionScaleCd          string    `bun:"dimension_scale_cd,type:char,default:('I')"`
}
