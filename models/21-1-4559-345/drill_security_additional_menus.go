package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type DrillSecurityAdditionalMenus struct {
	bun.BaseModel        `bun:"table:drill_security_additional_menus"`
	DsAdditionalMenusUid int32     `bun:"ds_additional_menus_uid,type:int,autoincrement,identity"`      // identity column
	BaseMenu             string    `bun:"base_menu,type:varchar(255)"`                                  // menu option most likely to exist
	DuplicateMenu        string    `bun:"duplicate_menu,type:varchar(255)"`                             // additional menu that shares same name
	WindowName           string    `bun:"window_name,type:varchar(255)"`                                // window name of menu option
	DateCreated          time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy            string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified     time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy     string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
