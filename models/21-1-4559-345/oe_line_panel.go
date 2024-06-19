package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type OeLinePanel struct {
	bun.BaseModel      `bun:"table:oe_line_panel"`
	OeLinePanelUid     int32     `bun:"oe_line_panel_uid,type:int,autoincrement,pk"`                  // Unique identifier for oe_line_panel
	DefaultName        string    `bun:"default_name,type:varchar(255),nullzero"`                      // Default name for this panel
	DefaultSequenceNo  int32     `bun:"default_sequence_no,type:int"`                                 // Default value for whether this is the default panel.
	DefaultDisplayInOe string    `bun:"default_display_in_oe,type:char(1)"`                           // Default value for whether this panel is displayed in oe.
	RealDw             string    `bun:"real_dw,type:char(1)"`                                         // Indicates if this is a real dw or user-created dw.
	SystemSettingName  string    `bun:"system_setting_name,type:varchar(128),nullzero"`               // Name of system setting associated with this panel.
	DateCreated        time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy          string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified   time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy   string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
