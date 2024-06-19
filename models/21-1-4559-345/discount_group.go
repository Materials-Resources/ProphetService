package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type DiscountGroup struct {
	bun.BaseModel            `bun:"table:discount_group"`
	DiscountGroupId          string    `bun:"discount_group_id,type:varchar(8),pk"`                      // What is the unique identifier for this discount group?
	DiscountGroupDescription string    `bun:"discount_group_description,type:varchar(40)"`               // How would you describe this discount group?
	DeleteFlag               string    `bun:"delete_flag,type:char(1)"`                                  // Indicates whether this record is logically deleted
	DateCreated              time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified         time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy         string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	ExtendedDesc             string    `bun:"extended_desc,type:varchar(255),nullzero"`                  // Extended description of the discount group.
}
