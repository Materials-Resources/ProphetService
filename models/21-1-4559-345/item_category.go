package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type ItemCategory struct {
	bun.BaseModel            `bun:"table:item_category"`
	ItemCategoryUid          int32     `bun:"item_category_uid,type:int,pk"`                                // Unique ID to identify Item Category
	ItemCategoryId           string    `bun:"item_category_id,type:varchar(255)"`                           // Short unique alphanumeric nameof the item category
	ItemCategoryDesc         string    `bun:"item_category_desc,type:varchar(255)"`                         // Long description of the item category
	MasterCategoryFlag       string    `bun:"master_category_flag,type:char(1),default:('N')"`              // Indicates whether this category contains items (vs. other categories)
	ParentCategoryFlag       string    `bun:"parent_category_flag,type:char(1),default:('N')"`              // Indicates that this category contains other categories
	DisplayOnWebFlag         string    `bun:"display_on_web_flag,type:char(1),default:('N')"`               // Indicates that the Item Category should display on the web for the user to drill into
	DisplayMasterProductFlag string    `bun:"display_master_product_flag,type:char(1),default:('N')"`       // Indicates that the images, descriptions, and links associated with the Master Product Category should be displayed on the web
	CatalogPage              string    `bun:"catalog_page,type:varchar(255),nullzero"`                      // Alphanumeric ID of the catalog page
	DeleteFlag               string    `bun:"delete_flag,type:char(1),default:('N')"`                       // Indicates that the item category has been deleted
	DateCreated              time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy                string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified         time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy         string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	SubCategoryImageFile     string    `bun:"sub_category_image_file,type:varchar(255),nullzero"`           // image file location for sub category
}
