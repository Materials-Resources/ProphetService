package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type PortalAssignment struct {
	bun.BaseModel       `bun:"table:portal_assignment"`
	PortalAssignmentUid int32     `bun:"portal_assignment_uid,type:int,pk"`                            // Unique ID for the portal_assignment
	PortalUid           int32     `bun:"portal_uid,type:int"`                                          // Unique ID for the portal
	UserId              string    `bun:"user_id,type:varchar(30),nullzero"`                            // User ID
	RoleUid             int32     `bun:"role_uid,type:int,nullzero"`                                   // Unique ID for the role
	DateCreated         time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy           string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified    time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy    string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	RowStatusFlag       int32     `bun:"row_status_flag,type:int,default:((704))"`
	EnabledForVersionCd int32     `bun:"enabled_for_version_cd,type:int,default:((1423))"` // Indicates whether portal is available for desktop, web or both
	RefreshTimer        int32     `bun:"refresh_timer,type:int,default:((0))"`             // time to take before refresh portal content
}
