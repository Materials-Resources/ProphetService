package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type PricingServiceMap struct {
	bun.BaseModel    `bun:"table:pricing_service_map"`
	LayoutId         float64   `bun:"layout_id,type:decimal(19,0),pk"`                           // Foreign key to pricing_service_layout_detail, identifies to which layout and column the rule applies
	LayoutDetailId   float64   `bun:"layout_detail_id,type:decimal(19,0)"`                       // Foreign key to pricing_service_layout_detail, identifies to which layout and column the rule applies
	TargetTable      string    `bun:"target_table,type:varchar(32),pk"`                          // The database table to be updated by the values for the associated column
	TargetColumn     string    `bun:"target_column,type:varchar(32),pk"`                         // The database column to be updated by the values for the associated column
	DeleteFlag       string    `bun:"delete_flag,type:char(1)"`                                  // Indicates whether this record is logically deleted
	DateCreated      time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
}
