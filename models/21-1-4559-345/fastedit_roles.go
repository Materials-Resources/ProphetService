package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type FasteditRoles struct {
	bun.BaseModel    `bun:"table:fastedit_roles"`
	FasteditRolesUid int32     `bun:"fastedit_roles_uid,type:int,autoincrement,pk"`                 // Identity column
	ClassName        string    `bun:"class_name,type:varchar(255)"`                                 // Class name of the window
	RoleUid          int32     `bun:"role_uid,type:int"`                                            // Id of corresponding role
	ActiveFlag       string    `bun:"active_flag,type:char(1),default:('Y')"`                       // Enable or disable fastedit for the role
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
