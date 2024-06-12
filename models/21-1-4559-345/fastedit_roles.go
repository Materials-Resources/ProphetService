package model

import (
	"github.com/uptrace/bun"
	"time"
)

type FasteditRoles struct {
	bun.BaseModel    `bun:"table:fastedit_roles"`
	FasteditRolesUid int32     `bun:"fastedit_roles_uid,type:int,pk,identity"`
	ClassName        string    `bun:"class_name,type:varchar(255)"`
	RoleUid          int32     `bun:"role_uid,type:int"`
	ActiveFlag       string    `bun:"active_flag,type:char,default:('Y')"`
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
