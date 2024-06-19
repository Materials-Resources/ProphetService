package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type InvPeriodUsage struct {
	bun.BaseModel               `bun:"table:inv_period_usage"`
	LocationId                  float64   `bun:"location_id,type:decimal(19,0),unique"`                     // Usage for which location?
	InvPeriodUsage              float64   `bun:"inv_period_usage,type:decimal(19,9)"`                       // Actual usage for the period.
	DateCreated                 time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified            time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy            string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	ScheduledUsage              float64   `bun:"scheduled_usage,type:decimal(19,9),nullzero"`               // What was the scheduled usage of material?
	ForecastUsage               float64   `bun:"forecast_usage,type:decimal(19,9),nullzero"`                // What is the forecast usage of this inventory item
	ForecastDeviationPercentage float64   `bun:"forecast_deviation_percentage,type:decimal(19,8),nullzero"` // Forecast deviation percentage.
	MadPercentage               float64   `bun:"mad_percentage,type:decimal(19,8),nullzero"`                // Mean average deviation percentage. Used for filtering usage.
	FilteredUsage               float64   `bun:"filtered_usage,type:decimal(19,9),nullzero"`                // The usage with aberrations filtered out.
	NumberOfOrders              float64   `bun:"number_of_orders,type:decimal(19,0)"`                       // The number of orders in the period.
	Edited                      string    `bun:"edited,type:char(1),nullzero"`                              // Indicates whether the filtered usage has been edited.
	NumberOfHits                float64   `bun:"number_of_hits,type:decimal(19,0),nullzero"`                // Number of orders fulfilled on-time.
	DemandPeriodUid             int32     `bun:"demand_period_uid,type:int,unique"`                         // Unique identifier for the demand period.
	InvMastUid                  int32     `bun:"inv_mast_uid,type:int,unique"`                              // Unique identifier for the item id.
	UsageCopied                 string    `bun:"usage_copied,type:char(1),default:('N')"`                   // Indicates that the usage was copied to a substitute item
	InvPeriodUsageUid           int32     `bun:"inv_period_usage_uid,type:int,autoincrement,identity,pk"`
	MeanAbsolutePercentError    float64   `bun:"mean_absolute_percent_error,type:decimal(19,4),nullzero"`        // Tracks accuracy of forecasting.  Used in advanced demand forecasting in determining forecasting error and erratic items.
	ExceptionalSalesFlag        string    `bun:"exceptional_sales_flag,type:char(1),nullzero"`                   // Flags whether filtered usage is being used as the actual usage because of exceptional sales.
	ExceptionalDeviationFlag    string    `bun:"exceptional_deviation_flag,type:char(1),nullzero"`               // Used in advanced demand forecasting.  Flags when the MAD Percentage to mean absolute percent error ratio is sufficiently large.
	ReviewedFlag                string    `bun:"reviewed_flag,type:char(1),nullzero"`                            // Indicates if this usage period was reviewed - for demand review window only
	LastReviewedDate            time.Time `bun:"last_reviewed_date,type:datetime,nullzero"`                      // Indicates last reviewed date of the usage period - for demand review window only
	UsageNotes                  string    `bun:"usage_notes,type:varchar(4000),nullzero"`                        // Notes on the item's usage for a period
	ForecastAdjustmentPercent   float64   `bun:"forecast_adjustment_percent,type:decimal(19,9),default:((100))"` // Percent that the forecast was adjusted by
	SavedFilteredUsage          float64   `bun:"saved_filtered_usage,type:decimal(19,9),nullzero"`               // The filtered usage will be stored here temporarily in cases when forecasting must temporarily overwrite the filtered usage for an item during calculation
	InvPeriodUsageThisLocation  float64   `bun:"inv_period_usage_this_location,type:decimal(19,9),nullzero"`
	Comments                    string    `bun:"comments,type:varchar(8000),nullzero"` // Allows user to make annotations. ex. why a filtered usage was edited
	Imported                    string    `bun:"imported,type:char(1),nullzero"`       // Whether record was imported; imported records are not affected by the rebuild
}
