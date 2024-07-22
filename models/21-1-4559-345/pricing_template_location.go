package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type PricingTemplateLocation struct {
	bun.BaseModel              `bun:"table:pricing_template_location"`
	PricingTemplateLocationUid int32     `bun:"pricing_template_location_uid,type:int,pk"`                    // Unique identifier for the record
	PricingTemplateUid         int32     `bun:"pricing_template_uid,type:int"`                                // Unique identifier for the associated pricing_template record
	LocationId                 float64   `bun:"location_id,type:decimal(19,0)"`                               // The location where values can be defaulted for.
	DateCreated                time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy                  string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified           time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy           string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
