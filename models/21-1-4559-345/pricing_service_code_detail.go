package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type PricingServiceCodeDetail struct {
	bun.BaseModel    `bun:"table:pricing_service_code_detail"`
	CodeId           float64   `bun:"code_id,type:decimal(19,0),pk"`                                 // FK to pricing_service_code.code_id
	CodeDetailId     float64   `bun:"code_detail_id,type:decimal(19,0),pk"`                          // System_assigned (sequential) value to uniquely identify detail record within code_id
	CodeValue        *string   `bun:"code_value,type:varchar(40)"`                                   // Code to be used in script
	ValueDesc        *string   `bun:"value_desc,type:varchar(254)"`                                  // Text to be displayed to user for code_value
	GroupCode        *string   `bun:"group_code,type:varchar(10)"`                                   // Code to group a set of codes which have the same code_id
	SubgroupCode     *string   `bun:"subgroup_code,type:varchar(10)"`                                // Code to group set of codes within group_code
	DeleteFlag       string    `bun:"delete_flag,type:char(1)"`                                      // Indicates whether this record is logically deleted
	DateCreated      time.Time `bun:"date_created,type:datetime"`                                    // Indicates the date/time this record was created.
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`                              // Indicates the date/time this record was last modified.
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"` // ID of the user who last maintained this record
}
