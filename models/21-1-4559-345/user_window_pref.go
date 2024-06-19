package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type UserWindowPref struct {
	bun.BaseModel     `bun:"table:user_window_pref"`
	UserWindowPrefUid int32     `bun:"user_window_pref_uid,type:int,autoincrement,identity,pk"`      // Unique identifier
	UserId            string    `bun:"user_id,type:varchar(30),unique"`                              // The user ID that these preferences are for
	WindowCd          int32     `bun:"window_cd,type:int,unique"`                                    // The window that these preferences are for
	ObjectName        string    `bun:"object_name,type:varchar(255),unique"`                         // The object on the window that we saved preferences for
	ObjectProperty    string    `bun:"object_property,type:varchar(255),unique"`                     // The property of the object we are saving preferences for
	ObjectValue       string    `bun:"object_value,type:varchar(8000)"`                              // The value or preference saved
	DateCreated       time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy         string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified  time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy  string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	WindowName        string    `bun:"window_name,type:varchar(255),unique,nullzero"`                // Class name of window with preference
}
