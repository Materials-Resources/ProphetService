package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type DcNavDrill struct {
	bun.BaseModel        `bun:"table:dc_nav_drill"`
	DcNavDrillUid        int32     `bun:"dc_nav_drill_uid,type:int,autoincrement,identity,pk"`          // Unique identifier for the table
	SourceWindow         string    `bun:"source_window,type:varchar(255)"`                              // Window drilling from
	SourceDatawindow     string    `bun:"source_datawindow,type:varchar(255),nullzero"`                 // DataWindow drilling from
	SourceField          string    `bun:"source_field,type:varchar(255)"`                               // Field that was clicked on
	DestWindow           string    `bun:"dest_window,type:varchar(255)"`                                // Window navigating to
	DestDatawindow       string    `bun:"dest_datawindow,type:varchar(255),nullzero"`                   // DataWindow navigating to
	DestField            string    `bun:"dest_field,type:varchar(255)"`                                 // Field to be populated on window open
	RowStatusFlag        int32     `bun:"row_status_flag,type:int,default:((704))"`                     // Status of row Active/Inactive/Delete
	DateCreated          time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy            string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified     time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy     string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	TypeCd               int32     `bun:"type_cd,type:int,default:((1418))"`                            // Code indicating whether a user-defined (1418) or standard (2005) drill down
	ApplyDrillToAllUsers string    `bun:"apply_drill_to_all_users,type:char(1),default:('Y')"`          // Indicates whether the drill down has global permissions
	DestWindowName       string    `bun:"dest_window_name,type:varchar(255),nullzero"`                  // Name of the drill target window
	SourceDataField      string    `bun:"source_data_field,type:varchar(255),nullzero"`                 // Used when data needed for drill is not in the clicked field.
	NavigationType       int32     `bun:"navigation_type,type:int,default:((3590))"`                    // Indicates whether record is a drill or a data pub/sub
}
