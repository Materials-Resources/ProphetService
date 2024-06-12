package model

type InvPeriodUsageTemp struct {
	bun.BaseModel         `bun:"table:inv_period_usage_temp"`
	LocationId            float64   `bun:"location_id,type:decimal(19,0)"`
	DemandPeriodUid       int32     `bun:"demand_period_uid,type:int"`
	InvMastUid            int32     `bun:"inv_mast_uid,type:int"`
	InvPeriodUsage        float64   `bun:"inv_period_usage,type:decimal(19,9)"`
	ScheduledUsage        float64   `bun:"scheduled_usage,type:decimal(19,9)"`
	NumberOfOrders        int32     `bun:"number_of_orders,type:int"`
	NumberOfHits          int32     `bun:"number_of_hits,type:int"`
	DateCreated           time.Time `bun:"date_created,type:datetime"`
	LastMaintainedBy      string    `bun:"last_maintained_by,type:varchar(255)"`
	InvPeriodUsageTempUid int32     `bun:"inv_period_usage_temp_uid,type:int,pk,identity"`
}
