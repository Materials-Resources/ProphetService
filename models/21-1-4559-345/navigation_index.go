package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type NavigationIndex struct {
	bun.BaseModel      `bun:"table:navigation_index"`
	NavigationIndexUid int32     `bun:"navigation_index_uid,type:int,autoincrement,identity,pk"`      // Unique ID for navigation index
	UsersId            string    `bun:"users_id,type:varchar(30),unique"`                             // User ID
	MenuName           string    `bun:"menu_name,type:varchar(2000),unique"`                          // Full hierarchical menu class name
	MenuText           string    `bun:"menu_text,type:varchar(255)"`                                  // Translated menu text
	RowStatusFlag      int32     `bun:"row_status_flag,type:int,default:((704))"`                     // Standard row status
	DateCreated        time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy          string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified   time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy   string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	WindowName         string    `bun:"window_name,type:varchar(255),nullzero"`                       // Stores the name of the window that gets opened by the menu item
}
