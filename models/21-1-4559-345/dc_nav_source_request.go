package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type DcNavSourceRequest struct {
	bun.BaseModel         `bun:"table:dc_nav_source_request"`
	DcNavSourceRequestUid int32     `bun:"dc_nav_source_request_uid,type:int,autoincrement,identity,pk"` // Unique identifier for the table
	SourceWindow          string    `bun:"source_window,type:varchar(255)"`                              // The source window for the request
	SourceDatawindow      *string   `bun:"source_datawindow,type:varchar(255)"`                          // The source datawindow for the request
	SourceField           string    `bun:"source_field,type:varchar(255)"`                               // The source field for the request
	DateCreated           time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy             string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified      time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy      string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	NavigationType        int32     `bun:"navigation_type,type:int,default:((3590))"`                    // Indicates whether record is a drill or a data pub/sub
}
