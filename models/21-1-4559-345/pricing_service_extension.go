package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type PricingServiceExtension struct {
	bun.BaseModel        `bun:"table:pricing_service_extension"`
	LayoutId             float64   `bun:"layout_id,type:decimal(19,0),pk"`                           // Foreign key to pricing_service_layout_detail, identifies to which layout and column the rule applies
	LayoutDetailId       float64   `bun:"layout_detail_id,type:decimal(19,0),pk"`                    // Foreign key to pricing_service_layout_detail, identifies to which layout and column the rule applies
	ExtensionId          float64   `bun:"extension_id,type:decimal(19,0),pk"`                        // Unique identifier within the associated pricing service layout definition and column
	ExtensionSource      *string   `bun:"extension_source,type:varchar(10)"`                         // Code to identify the source of the conversion value: another column or manually entered
	ExtensionValue       *string   `bun:"extension_value,type:varchar(50)"`                          // Manually entered value for the extension rule
	SourceLayoutId       *float64  `bun:"source_layout_id,type:decimal(19,0)"`                       // Identifies the source column when the factor value source is from another column
	SourceLayoutDetailId *float64  `bun:"source_layout_detail_id,type:decimal(19,0)"`                // Identifies the source column when the factor value source is from another column
	SequenceNumber       *float64  `bun:"sequence_number,type:decimal(19,0)"`                        // Indicates the sequence in which to process multiple extension rules for a column
	DeleteFlag           string    `bun:"delete_flag,type:char(1)"`                                  // Indicates whether this record is logically deleted
	DateCreated          time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified     time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy     string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	ExtensionPosition    *string   `bun:"extension_position,type:varchar(10)"`                       // This column will allow the user to indicate whethe
	SourceStart          *int32    `bun:"source_start,type:int"`                                     // To allow user to select PART of another column. (e
	SourceLength         *int32    `bun:"source_length,type:int"`                                    // To allow user to select PART of another column. (e
}
