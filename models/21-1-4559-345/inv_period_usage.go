package model

import (
	"github.com/uptrace/bun"
	"time"
)

type InvPeriodUsage struct {
	bun.BaseModel               `bun:"table:inv_period_usage"`
	LocationId                  float64   `bun:"location_id,type:decimal(19,0)"`
	InvPeriodUsage              float64   `bun:"inv_period_usage,type:decimal(19,9)"`
	DateCreated                 time.Time `bun:"date_created,type:datetime"`
	DateLastModified            time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy            string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	ScheduledUsage              float64   `bun:"scheduled_usage,type:decimal(19,9),nullzero"`
	ForecastUsage               float64   `bun:"forecast_usage,type:decimal(19,9),nullzero"`
	ForecastDeviationPercentage float64   `bun:"forecast_deviation_percentage,type:decimal(19,8),nullzero"`
	MadPercentage               float64   `bun:"mad_percentage,type:decimal(19,8),nullzero"`
	FilteredUsage               float64   `bun:"filtered_usage,type:decimal(19,9),nullzero"`
	NumberOfOrders              float64   `bun:"number_of_orders,type:decimal(19,0)"`
	Edited                      string    `bun:"edited,type:char,nullzero"`
	NumberOfHits                float64   `bun:"number_of_hits,type:decimal(19,0),nullzero"`
	DemandPeriodUid             int32     `bun:"demand_period_uid,type:int"`
	InvMastUid                  int32     `bun:"inv_mast_uid,type:int"`
	UsageCopied                 string    `bun:"usage_copied,type:char,default:('N')"`
	InvPeriodUsageUid           int32     `bun:"inv_period_usage_uid,type:int,pk,identity"`
	MeanAbsolutePercentError    float64   `bun:"mean_absolute_percent_error,type:decimal(19,4),nullzero"`
	ExceptionalSalesFlag        string    `bun:"exceptional_sales_flag,type:char,nullzero"`
	ExceptionalDeviationFlag    string    `bun:"exceptional_deviation_flag,type:char,nullzero"`
	ReviewedFlag                string    `bun:"reviewed_flag,type:char,nullzero"`
	LastReviewedDate            time.Time `bun:"last_reviewed_date,type:datetime,nullzero"`
	UsageNotes                  string    `bun:"usage_notes,type:varchar(4000),nullzero"`
	ForecastAdjustmentPercent   float64   `bun:"forecast_adjustment_percent,type:decimal(19,9),default:((100))"`
	SavedFilteredUsage          float64   `bun:"saved_filtered_usage,type:decimal(19,9),nullzero"`
	InvPeriodUsageThisLocation  float64   `bun:"inv_period_usage_this_location,type:decimal(19,9),nullzero"`
	Comments                    string    `bun:"comments,type:varchar(8000),nullzero"`
	Imported                    string    `bun:"imported,type:char,nullzero"`
}
