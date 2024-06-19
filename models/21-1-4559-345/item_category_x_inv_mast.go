package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type ItemCategoryXInvMast struct {
	bun.BaseModel           `bun:"table:item_category_x_inv_mast"`
	ItemCategoryXInvMastUid int32     `bun:"item_category_x_inv_mast_uid,type:int,pk"`                     // Unique ID to identify Item Category / Item link
	ItemCategoryUid         int32     `bun:"item_category_uid,type:int,unique"`                            // Item Category (UID) with which the item is associated
	InvMastUid              int32     `bun:"inv_mast_uid,type:int,unique"`                                 // Item associated with the item category
	SequenceNo              int32     `bun:"sequence_no,type:int"`                                         // Sequence of this item in the category
	DisplayOnWebFlag        string    `bun:"display_on_web_flag,type:char(1),default:('N')"`               // Indicates whether this item should be displayed on the web as part of the category
	DeleteFlag              string    `bun:"delete_flag,type:char(1),default:('N')"`                       // Indicates that the link has been deleted
	DateCreated             time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy               string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	DisplayDesc             string    `bun:"display_desc,type:varchar(255)"`                               // Item description - overrides inv_mast.item_desc in category context
	Comments                string    `bun:"comments,type:varchar(255),nullzero"`                          // Additional user comments.
	AlternateCodeUid        int32     `bun:"alternate_code_uid,type:int,nullzero"`                         // Unique identifier for an alternate code.
}
