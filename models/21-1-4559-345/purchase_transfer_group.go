package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type PurchaseTransferGroup struct {
	bun.BaseModel             `bun:"table:purchase_transfer_group"`
	PurchaseTransferGroupId   string    `bun:"purchase_transfer_group_id,type:varchar(8),pk"`                 // Unique ID for the Group (User defined)
	PurchaseTransferGroupDesc string    `bun:"purchase_transfer_group_desc,type:varchar(40)"`                 // Description of the purchase transfer group.
	DeleteFlag                string    `bun:"delete_flag,type:char(1)"`                                      // Indicates whether this record is logically deleted
	DateCreated               time.Time `bun:"date_created,type:datetime"`                                    // Indicates the date/time this record was created.
	DateLastModified          time.Time `bun:"date_last_modified,type:datetime"`                              // Indicates the date/time this record was last modified.
	LastMaintainedBy          string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"` // ID of the user who last maintained this record
	RegionGroupType           string    `bun:"region_group_type,type:char(1),default:('N')"`                  // Indicates if this group is a Region Group (value R) or Sister Site Group (value S) or neither (value N).
}
