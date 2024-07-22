package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type DwobjectDependency struct {
	bun.BaseModel         `bun:"table:dwobject_dependency"`
	DwobjectDependencyUid int32     `bun:"dwobject_dependency_uid,type:int,autoincrement,identity,pk"`   // Unique ID for this record.
	PrimaryObjectName     string    `bun:"primary_object_name,type:varchar(255)"`                        // Primary DW object that other DWs are dependent on.
	SecondaryObjectName   string    `bun:"secondary_object_name,type:varchar(255)"`                      // DW object that must be modified when the primary DW object is modified.
	DateCreated           time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy             string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified      time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy      string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
