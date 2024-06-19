package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type ExtensibilityWindow struct {
	bun.BaseModel          `bun:"table:extensibility_window"`
	ExtensibilityWindowUid int32     `bun:"extensibility_window_uid,type:int,autoincrement,scanonly,pk"`  // UID for the table.
	WindowName             string    `bun:"window_name,type:varchar(255)"`                                // Window class name that the extensibility functionality will be enalbed for.
	DateCreated            time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy              string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified       time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy       string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	EnableFlag             string    `bun:"enable_flag,type:char(1),default:('Y')"`                       // Whether window is enabled or disabled for extensibility
}
