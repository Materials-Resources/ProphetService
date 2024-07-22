package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type Module struct {
	bun.BaseModel     `bun:"table:module"`
	ModuleId          string    `bun:"module_id,type:varchar(8),pk"`                                 // What is the unique identifier for this module?
	ModuleDescription string    `bun:"module_description,type:varchar(40)"`                          // What is this module for?
	FrameName         *string   `bun:"frame_name,type:varchar(255)"`                                 // Holds the frame window class for the module.  Used in opening the app with a command line parm.
	ClassName         *string   `bun:"class_name,type:varchar(255)"`                                 // Menu class from the frame_menu table
	ModuleGroupNo     *int32    `bun:"module_group_no,type:int"`                                     // Major menu group for the module
	DateCreated       time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy         string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified  time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy  string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
