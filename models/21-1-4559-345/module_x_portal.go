package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type ModuleXPortal struct {
	bun.BaseModel    `bun:"table:module_x_portal"`
	ModuleXPortalUid int32     `bun:"module_x_portal_uid,type:int,autoincrement,scanonly,pk"`       // Unique identifier
	ModuleCd         int32     `bun:"module_cd,type:int"`                                           // Identifies the module
	PortalCd         int32     `bun:"portal_cd,type:int"`                                           // Identifies the portal within the module
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
