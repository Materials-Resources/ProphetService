package model

import (
	"time"

	"github.com/uptrace/bun"
)

type AverageInventoryValue struct {
	bun.BaseModel            `bun:"table:average_inventory_value"`
	AverageInventoryValueUid int32     `bun:"average_inventory_value_uid,type:int,pk,identity"`
	DemandPeriodUid          int32     `bun:"demand_period_uid,type:int"`
	LocationId               float64   `bun:"location_id,type:decimal(19,0)"`
	InvMastUid               int32     `bun:"inv_mast_uid,type:int"`
	InventoryValue           float64   `bun:"inventory_value,type:decimal(19,9)"`
	NoOfUpdates              int32     `bun:"no_of_updates,type:int"`
	DateCreated              time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified         time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy         string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
