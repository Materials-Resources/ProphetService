package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type ItemCategoryHierarchyDeleteAudit struct {
	bun.BaseModel                       `bun:"table:item_category_hierarchy_delete_audit"`
	ItemCategoryHierarchyDeleteAuditUid int32     `bun:"item_category_hierarchy_delete_audit_uid,type:int,autoincrement,identity"`
	ParentItemCategoryUid               int32     `bun:"parent_item_category_uid,type:int"`
	ChildItemCategoryUid                int32     `bun:"child_item_category_uid,type:int"`
	ParentItemCategoryId                string    `bun:"parent_item_category_id,type:varchar(255)"`
	ChildItemCategoryId                 string    `bun:"child_item_category_id,type:varchar(255)"`
	DateDeleted                         time.Time `bun:"date_deleted,type:datetime"`
	UserDeleted                         *string   `bun:"user_deleted,type:Error: 50000, Severity: -1, State: 1. (Params:). The error is printed in terse mode because there was error during formatting. Tracing, ETW, notifications etc are skipped.\r\n"`
	ItemCategoryPath                    *string   `bun:"item_category_path,type:varchar(max)"`
	ItemCategoryDesc                    *string   `bun:"item_category_desc,type:varchar(max)"`
	CatLevel                            *string   `bun:"cat_level,type:varchar(max)"`
}
