package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type ItemCategoryHierarchy struct {
	bun.BaseModel            `bun:"table:item_category_hierarchy"`
	ItemCategoryHierarchyUid int32     `bun:"item_category_hierarchy_uid,type:int,pk"`                      // Unique ID for Item Category hierarchy
	ParentItemCategoryUid    int32     `bun:"parent_item_category_uid,type:int"`                            // Unique ID for Parent Item Category
	ChildItemCategoryUid     int32     `bun:"child_item_category_uid,type:int"`                             // Unique ID for Child Item Category
	DateCreated              time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy                string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified         time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy         string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	SequenceNo               int32     `bun:"sequence_no,type:int,default:(0)"`
}
