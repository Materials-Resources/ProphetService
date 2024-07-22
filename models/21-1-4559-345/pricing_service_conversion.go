package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type PricingServiceConversion struct {
	bun.BaseModel        `bun:"table:pricing_service_conversion"`
	LayoutId             float64   `bun:"layout_id,type:decimal(19,0),pk"`                           // Foreign key to pricing_service_layout_detail, identifies to which layout and column the rule applies
	LayoutDetailId       float64   `bun:"layout_detail_id,type:decimal(19,0),pk"`                    // Foreign key to pricing_service_layout_detail, identifies to which layout and column the rule applies
	ConversionId         float64   `bun:"conversion_id,type:decimal(19,0),pk"`                       // Unique identifier within the associated pricing service layout definition and column
	SequenceNumber       *float64  `bun:"sequence_number,type:decimal(19,0)"`                        // Indicates the sequence in which to process multiple conversion rules for a column
	FactorSource         *string   `bun:"factor_source,type:varchar(10)"`                            // Code to identify the source of the factor value: another column or manually entered
	SourceLayoutId       *float64  `bun:"source_layout_id,type:decimal(19,0)"`                       // Identifies the source column when the factor value source is from another column
	SourceLayoutDetailId *float64  `bun:"source_layout_detail_id,type:decimal(19,0)"`                // Identifies the source column when the factor value source is from another column
	Factor               *float64  `bun:"factor,type:decimal(19,4)"`                                 // The value to be applied to the base value by multiplication, division, addition, subtraction
	Operand              *string   `bun:"operand,type:varchar(10)"`                                  // The arithmetic operation to be performed on the base value using the factor
	DeleteFlag           string    `bun:"delete_flag,type:char(1)"`                                  // Indicates whether this record is logically deleted
	DateCreated          time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified     time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy     string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
}
