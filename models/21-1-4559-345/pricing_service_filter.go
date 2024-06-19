package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type PricingServiceFilter struct {
	bun.BaseModel    `bun:"table:pricing_service_filter"`
	LayoutId         float64   `bun:"layout_id,type:decimal(19,0),pk"`                           // Foreign key to pricing_service_layout_detail, identifies to which layout and column the rule applies
	LayoutDetailId   float64   `bun:"layout_detail_id,type:decimal(19,0),pk"`                    // Foreign key to pricing_service_layout_detail, identifies to which layout and column the rule applies
	FilterId         float64   `bun:"filter_id,type:decimal(19,0),pk"`                           // Unique identifier within the associated pricing service layout definition and column
	Operand          string    `bun:"operand,type:varchar(255),nullzero"`                        // The comparison operator for the filter
	FilterValue      string    `bun:"filter_value,type:varchar(50),nullzero"`                    // The value to be used as the comparison in the filter
	DeleteFlag       string    `bun:"delete_flag,type:char(1)"`                                  // Indicates whether this record is logically deleted
	DateCreated      time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
}
