package model

import (
	"github.com/uptrace/bun"
	"time"
)

type ItemCategoryHierarchy struct {
	bun.BaseModel            `bun:"table:item_category_hierarchy"`
	ItemCategoryHierarchyUid int32     `bun:"item_category_hierarchy_uid,type:int,pk"`
	ParentItemCategoryUid    int32     `bun:"parent_item_category_uid,type:int"`
	ChildItemCategoryUid     int32     `bun:"child_item_category_uid,type:int"`
	DateCreated              time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified         time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy         string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	SequenceNo               int32     `bun:"sequence_no,type:int,default:(0)"`
}
