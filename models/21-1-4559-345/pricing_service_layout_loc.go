package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type PricingServiceLayoutLoc struct {
	bun.BaseModel              `bun:"table:pricing_service_layout_loc"`
	PricingServiceLayoutLocUid int32     `bun:"pricing_service_layout_loc_uid,type:int,autoincrement,identity,pk"` // Unique identifier for the record
	LayoutId                   float64   `bun:"layout_id,type:decimal(19,0)"`                                      // The unique identifier for the Pricing Service Layout
	LocationId                 float64   `bun:"location_id,type:decimal(19,0)"`                                    // The location where the items will be built/updated.
	DateCreated                time.Time `bun:"date_created,type:datetime,default:(getdate())"`                    // Date and time the record was originally created
	CreatedBy                  string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`              // User who created the record
	DateLastModified           time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`              // Date and time the record was modified
	LastMaintainedBy           string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`      // User who last changed the record
}
