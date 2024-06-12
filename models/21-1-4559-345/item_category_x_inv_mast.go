package model

import (
	"github.com/uptrace/bun"
	"time"
)

type ItemCategoryXInvMast struct {
	bun.BaseModel           `bun:"table:item_category_x_inv_mast"`
	ItemCategoryXInvMastUid int32     `bun:"item_category_x_inv_mast_uid,type:int,pk"`
	ItemCategoryUid         int32     `bun:"item_category_uid,type:int"`
	InvMastUid              int32     `bun:"inv_mast_uid,type:int"`
	SequenceNo              int32     `bun:"sequence_no,type:int"`
	DisplayOnWebFlag        string    `bun:"display_on_web_flag,type:char,default:('N')"`
	DeleteFlag              string    `bun:"delete_flag,type:char,default:('N')"`
	DateCreated             time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy               string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	DisplayDesc             string    `bun:"display_desc,type:varchar(255)"`
	Comments                string    `bun:"comments,type:varchar(255),nullzero"`
	AlternateCodeUid        int32     `bun:"alternate_code_uid,type:int,nullzero"`
}
