package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type UsersXOeLinePanel struct {
	bun.BaseModel        `bun:"table:users_x_oe_line_panel"`
	UsersXOeLinePanelUid int32     `bun:"users_x_oe_line_panel_uid,type:int,autoincrement,scanonly,pk"` // Unique identifier for users_x_oe_line_panel
	UsersId              string    `bun:"users_id,type:varchar(30),unique"`                             // id from the users table.
	OeLinePanelUid       int32     `bun:"oe_line_panel_uid,type:int,unique"`                            // oe_line_panel_uid from oe_line_panel table.
	PanelName            string    `bun:"panel_name,type:varchar(40),nullzero"`                         // The name of this panel.
	SequenceNo           int32     `bun:"sequence_no,type:int"`                                         // The sequence number for this panel/user.
	DateCreated          time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy            string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified     time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy     string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
