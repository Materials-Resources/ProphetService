package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type DcNavDrillXRoles struct {
	bun.BaseModel       `bun:"table:dc_nav_drill_x_roles"`
	DcNavDrillXRolesUid int32     `bun:"dc_nav_drill_x_roles_uid,type:int,autoincrement,pk"`           // Unique identifier for the table
	DcNavDrillUid       int32     `bun:"dc_nav_drill_uid,type:int"`                                    // Identifier for the dc_nav_drill_table
	RoleUid             int32     `bun:"role_uid,type:int"`                                            // Identifier for the roles table
	DateCreated         time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy           string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified    time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy    string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
