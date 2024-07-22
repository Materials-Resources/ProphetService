package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type FrameMenuReporting struct {
	bun.BaseModel         `bun:"table:frame_menu_reporting"`
	FrameMenuReportingUid int32     `bun:"frame_menu_reporting_uid,type:int,autoincrement,identity,pk"`  // Unique identifier
	DateCreated           time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy             string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified      time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy      string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	FrameMenuUid          int32     `bun:"frame_menu_uid,type:int"`                                      // Unique identifier for the frame_menu table
	ReportSyntax          *string   `bun:"report_syntax,type:varchar(max)"`                              // The meta data used to render the report
}
