package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type AverageInventoryValue struct {
	bun.BaseModel            `bun:"table:average_inventory_value"`
	AverageInventoryValueUid int32     `bun:"average_inventory_value_uid,type:int,autoincrement,scanonly,pk"` // Unique identifier for the record
	DemandPeriodUid          int32     `bun:"demand_period_uid,type:int"`                                     // Unique identifier for the demand period when the inventory value was calculated
	LocationId               float64   `bun:"location_id,type:decimal(19,0)"`                                 // The location where the item is located
	InvMastUid               int32     `bun:"inv_mast_uid,type:int"`                                          // Unique identifier for the item id specific to this record.
	InventoryValue           float64   `bun:"inventory_value,type:decimal(19,9)"`                             // The average inventory value for them item/location during the current demand period
	NoOfUpdates              int32     `bun:"no_of_updates,type:int"`                                         // Number of times the record has been updated during the demand period
	DateCreated              time.Time `bun:"date_created,type:datetime,default:(getdate())"`                 // Date and time the record was originally created
	CreatedBy                string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`           // User who created the record
	DateLastModified         time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`           // Date and time the record was modified
	LastMaintainedBy         string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`   // User who last changed the record
}
