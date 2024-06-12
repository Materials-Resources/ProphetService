package model

import (
	"github.com/uptrace/bun"
	"time"
)

type WindowXMenu struct {
	bun.BaseModel    `bun:"table:window_x_menu"`
	WindowXMenuUid   int32     `bun:"window_x_menu_uid,type:int,pk,identity"`
	WindowName       string    `bun:"window_name,type:varchar(255)"`
	MenuName         string    `bun:"menu_name,type:varchar(2000)"`
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
