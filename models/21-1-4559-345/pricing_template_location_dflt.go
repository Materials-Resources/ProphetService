package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type PricingTemplateLocationDflt struct {
	bun.BaseModel                  `bun:"table:pricing_template_location_dflt"`
	PricingTemplateLocationDfltUid int32     `bun:"pricing_template_location_dflt_uid,type:int,autoincrement,identity,pk"` // Unique identifier for the record
	PricingTemplateLocationUid     int32     `bun:"pricing_template_location_uid,type:int"`                                // Unique identifier for the associated pricing_template_location record
	PricingTemplateKeyFieldUid     int32     `bun:"pricing_template_key_field_uid,type:int"`                               // Unique identifier for the associated pricing_template_key_field record
	ColumnId                       float64   `bun:"column_id,type:decimal(19,0)"`                                          // Unqiue identifier for the pricing_service_column record
	DefaultValue                   string    `bun:"default_value,type:varchar(50)"`                                        // The default value for the column
	DateCreated                    time.Time `bun:"date_created,type:datetime,default:(getdate())"`                        // Date and time the record was originally created
	CreatedBy                      string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`                  // User who created the record
	DateLastModified               time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`                  // Date and time the record was modified
	LastMaintainedBy               string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`          // User who last changed the record
}
