package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type InvSub struct {
	bun.BaseModel    `bun:"table:inv_sub"`
	InvMastUid       int32     `bun:"inv_mast_uid,type:int,pk"`                                  // Unique identifier for the item id.
	SubInvMastUid    int32     `bun:"sub_inv_mast_uid,type:int,pk"`                              // Unique Identifier for the substitute item.
	SubDesc          string    `bun:"sub_desc,type:varchar(40)"`                                 // How would you describe this substitution?
	DeleteFlag       string    `bun:"delete_flag,type:char(1)"`                                  // Indicates whether this record is logically deleted
	DateCreated      time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	Interchangeable  string    `bun:"interchangeable,type:char(1)"`                              // Are the item and the substitute interchangeable?
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
}
