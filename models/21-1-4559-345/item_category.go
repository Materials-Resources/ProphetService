package model

import (
	"github.com/uptrace/bun"
	"time"
)

type ItemCategory struct {
	bun.BaseModel            `bun:"table:item_category"`
	ItemCategoryUid          int32     `bun:"item_category_uid,type:int,pk"`
	ItemCategoryId           string    `bun:"item_category_id,type:varchar(255)"`
	ItemCategoryDesc         string    `bun:"item_category_desc,type:varchar(255)"`
	MasterCategoryFlag       string    `bun:"master_category_flag,type:char,default:('N')"`
	ParentCategoryFlag       string    `bun:"parent_category_flag,type:char,default:('N')"`
	DisplayOnWebFlag         string    `bun:"display_on_web_flag,type:char,default:('N')"`
	DisplayMasterProductFlag string    `bun:"display_master_product_flag,type:char,default:('N')"`
	CatalogPage              string    `bun:"catalog_page,type:varchar(255),nullzero"`
	DeleteFlag               string    `bun:"delete_flag,type:char,default:('N')"`
	DateCreated              time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified         time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy         string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	SubCategoryImageFile     string    `bun:"sub_category_image_file,type:varchar(255),nullzero"`
}
