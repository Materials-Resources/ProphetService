package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type InvPeriodUsageTemp struct {
	bun.BaseModel         `bun:"table:inv_period_usage_temp"`
	LocationId            float64   `bun:"location_id,type:decimal(19,0)"`                               // The location ID for which usage data must be updated.
	DemandPeriodUid       int32     `bun:"demand_period_uid,type:int"`                                   // The demand_period_uid for which usage data must be updated.
	InvMastUid            int32     `bun:"inv_mast_uid,type:int"`                                        // Unique identifier for the item id.
	InvPeriodUsage        float64   `bun:"inv_period_usage,type:decimal(19,9)"`                          // The quantity that must be added to usage.
	ScheduledUsage        float64   `bun:"scheduled_usage,type:decimal(19,9)"`                           // The quantity that must be added to schedule usage.
	NumberOfOrders        int32     `bun:"number_of_orders,type:int"`                                    // The quantity that must be added to the number of orders.
	NumberOfHits          int32     `bun:"number_of_hits,type:int"`                                      // The quantity that must be added to the number of hits.
	DateCreated           time.Time `bun:"date_created,type:datetime"`                                   // Indicates the date/time this record was created.
	LastMaintainedBy      string    `bun:"last_maintained_by,type:varchar(255)"`                         // ID of the user who last maintained this record
	InvPeriodUsageTempUid int32     `bun:"inv_period_usage_temp_uid,type:int,autoincrement,identity,pk"` // Unique ID for inv_period_usage_temp records
}
