package model

type ItemCategoryHierarchyDeleteAudit struct {
	bun.BaseModel                       `bun:"table:item_category_hierarchy_delete_audit"`
	ItemCategoryHierarchyDeleteAuditUid int32     `bun:"item_category_hierarchy_delete_audit_uid,type:int,identity"`
	ParentItemCategoryUid               int32     `bun:"parent_item_category_uid,type:int"`
	ChildItemCategoryUid                int32     `bun:"child_item_category_uid,type:int"`
	ParentItemCategoryId                string    `bun:"parent_item_category_id,type:varchar(255)"`
	ChildItemCategoryId                 string    `bun:"child_item_category_id,type:varchar(255)"`
	DateDeleted                         time.Time `bun:"date_deleted,type:datetime"`
	UserDeleted                         `bun:"user_deleted,type:nvarchar(256),nullzero"`
	ItemCategoryPath                    string `bun:"item_category_path,type:varchar,nullzero"`
	ItemCategoryDesc                    string `bun:"item_category_desc,type:varchar,nullzero"`
	CatLevel                            string `bun:"cat_level,type:varchar,nullzero"`
}
