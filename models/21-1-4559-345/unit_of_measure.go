package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type UnitOfMeasure struct {
	bun.BaseModel         `bun:"table:unit_of_measure"`
	UnitId                string    `bun:"unit_id,type:varchar(8),pk"`                                // What is the unique identifier for this unit of measure?
	UnitDescription       string    `bun:"unit_description,type:varchar(30)"`                         // What is this unit of measure for?
	DeleteFlag            string    `bun:"delete_flag,type:char(1)"`                                  // Indicates whether this record is logically deleted
	DateCreated           time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified      time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy      string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	DisplayDescription    *string   `bun:"display_description,type:varchar(30)"`                      // What is the description of the unit of measure?
	RepackagingCostFactor *float64  `bun:"repackaging_cost_factor,type:decimal(19,9)"`                // A dollar amount of the cost for repackaging items
	MaritimeUom           *string   `bun:"maritime_uom,type:varchar(255)"`                            // Need to translate the maritime standard UOM to the P21 UOM
	PackagingUnitFlag     string    `bun:"packaging_unit_flag,type:char(1),default:('N')"`            // Identifies the UOM as a packaging unit
	DimensionScale        *string   `bun:"dimension_scale,type:varchar(2)"`                           // Track dimension scale for UOM
}
