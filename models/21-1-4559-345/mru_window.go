package model

import (
	"github.com/uptrace/bun"
	"time"
)

type MruWindow struct {
	bun.BaseModel    `bun:"table:mru_window"`
	MruWindowUid     int32     `bun:"mru_window_uid,type:int,pk"`
	WindowName       string    `bun:"window_name,type:varchar(255)"`
	UserId           string    `bun:"user_id,type:varchar(30)"`
	LineNo           int32     `bun:"line_no,type:int"`
	FrameName        string    `bun:"frame_name,type:varchar(255)"`
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	OpenedbyMenuitem string    `bun:"openedby_menuitem,type:varchar(128)"`
	Menuname         string    `bun:"menuname,type:varchar(128)"`
}
