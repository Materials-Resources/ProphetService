package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type WindowXMenu struct {
	bun.BaseModel    `bun:"table:window_x_menu"`
	WindowXMenuUid   int32     `bun:"window_x_menu_uid,type:int,autoincrement,scanonly,pk"`         // Unique identifier for the table
	WindowName       string    `bun:"window_name,type:varchar(255)"`                                // Window to be linked to the menu item
	MenuName         string    `bun:"menu_name,type:varchar(2000)"`                                 // Menu item captured by the navigation index
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
