package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type PricingServiceValueDetail struct {
	bun.BaseModel    `bun:"table:pricing_service_value_detail"`
	ValueId          float64   `bun:"value_id,type:decimal(19,0),pk"`                            // Foreign key to pricing_service_value, identifies to which value list the codes apply
	ValueDetailId    float64   `bun:"value_detail_id,type:decimal(19,0),pk"`                     // Unique identifier within value_id
	OriginalValue    string    `bun:"original_value,type:varchar(50),nullzero"`                  // The value from the import file to match on
	ConvertedValue   string    `bun:"converted_value,type:varchar(50),nullzero"`                 // The value to change the import file value to when a match is found
	DeleteFlag       string    `bun:"delete_flag,type:char(1)"`                                  // Indicates whether this record is logically deleted
	DateCreated      time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
}
