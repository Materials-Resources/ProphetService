package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type MruWindow struct {
	bun.BaseModel    `bun:"table:mru_window"`
	MruWindowUid     int32     `bun:"mru_window_uid,type:int,pk"`                                   // Unique identifier for mru_window
	WindowName       string    `bun:"window_name,type:varchar(255)"`                                // Recently used window name
	UserId           string    `bun:"user_id,type:varchar(30),unique"`                              // Which user owns this record
	LineNo           int32     `bun:"line_no,type:int,unique"`                                      // Used to track which window is more recently used
	FrameName        string    `bun:"frame_name,type:varchar(255),unique"`                          // Frame for which this record is for
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	OpenedbyMenuitem string    `bun:"openedby_menuitem,type:varchar(128)"`
	Menuname         string    `bun:"menuname,type:varchar(128)"`
}
