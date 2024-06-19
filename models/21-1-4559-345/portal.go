package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type Portal struct {
	bun.BaseModel    `bun:"table:portal"`
	PortalUid        int32     `bun:"portal_uid,type:int,pk"`                                       // Unique ID for the portal
	PortalName       string    `bun:"portal_name,type:varchar(255),unique"`                         // Short name for the portal
	PortalLayout     int32     `bun:"portal_layout,type:int"`                                       // Layout index for the portal
	PortalDesc       string    `bun:"portal_desc,type:varchar(255),nullzero"`                       // Description
	RowStatusFlag    int32     `bun:"row_status_flag,type:int,default:((704))"`                     // Indicates status of record (active/delete)
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	RefreshTimer     int32     `bun:"refresh_timer,type:int,default:((0))"`                         // time to take before refresh portal content
}
