package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type DrillSecurity struct {
	bun.BaseModel    `bun:"table:drill_security"`
	UsersId          string    `bun:"users_id,type:varchar(30),unique"`                             // User ID
	MenuName         string    `bun:"menu_name,type:varchar(500),unique"`                           // Menu object name
	WindowName       string    `bun:"window_name,type:varchar(255),nullzero"`                       // Window object name
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
