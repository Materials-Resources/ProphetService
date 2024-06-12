package model

import (
	"github.com/uptrace/bun"
	"time"
)

type UsersXOeLinePanel struct {
	bun.BaseModel        `bun:"table:users_x_oe_line_panel"`
	UsersXOeLinePanelUid int32     `bun:"users_x_oe_line_panel_uid,type:int,pk,identity"`
	UsersId              string    `bun:"users_id,type:varchar(30)"`
	OeLinePanelUid       int32     `bun:"oe_line_panel_uid,type:int"`
	PanelName            string    `bun:"panel_name,type:varchar(40),nullzero"`
	SequenceNo           int32     `bun:"sequence_no,type:int"`
	DateCreated          time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy            string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified     time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy     string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
