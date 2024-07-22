package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type PortalParamValue struct {
	bun.BaseModel       `bun:"table:portal_param_value"`
	PortalParamValueUid int32     `bun:"portal_param_value_uid,type:int,autoincrement,identity,pk"`    // Unique ID for the portal parameter value
	PortalAssignmentUid int32     `bun:"portal_assignment_uid,type:int,unique"`                        // Unique ID for the portal assignment
	PortalParamDefUid   int32     `bun:"portal_param_def_uid,type:int,unique"`                         // Unique ID for the portal element parameter definition
	Value               string    `bun:"value,type:varchar(255)"`                                      // Value to be passed to the portal for this user/role
	DateCreated         time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy           string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified    time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy    string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
