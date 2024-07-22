package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type PricingServiceValue struct {
	bun.BaseModel    `bun:"table:pricing_service_value"`
	ValueId          float64   `bun:"value_id,type:decimal(19,0),pk"`                            // Unique identifier
	ValueDesc        *string   `bun:"value_desc,type:varchar(50)"`                               // User entered description of the code set
	DeleteFlag       string    `bun:"delete_flag,type:char(1)"`                                  // Indicates whether this record is logically deleted
	DateCreated      time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	UseExactMatch    *string   `bun:"use_exact_match,type:char(1),default:('Y')"`                // Indicates whether the import file values match value exactly, or by like comparison
	ProcessType      *string   `bun:"process_type,type:varchar(10),default:('PSDESIGN')"`        // Type of process the list supports - Pricing Service vs. XML Conversion
}
