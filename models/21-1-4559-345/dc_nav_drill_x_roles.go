package model

import (
	"github.com/uptrace/bun"
	"time"
)

type DcNavDrillXRoles struct {
	bun.BaseModel       `bun:"table:dc_nav_drill_x_roles"`
	DcNavDrillXRolesUid int32     `bun:"dc_nav_drill_x_roles_uid,type:int,pk,identity"`
	DcNavDrillUid       int32     `bun:"dc_nav_drill_uid,type:int"`
	RoleUid             int32     `bun:"role_uid,type:int"`
	DateCreated         time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy           string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified    time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy    string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
