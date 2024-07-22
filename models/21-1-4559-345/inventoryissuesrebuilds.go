package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type Inventoryissuesrebuilds struct {
	bun.BaseModel    `bun:"table:inventoryissuesrebuilds"`
	Rebuilduid       int32     `bun:"rebuilduid,type:int,autoincrement,identity,pk"`                // Key value for this rebuild
	RebuildSp        string    `bun:"rebuild_sp,type:varchar(255)"`                                 // Stored procedure that will perform the rebuild.
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	RebuildEnabled   *bool     `bun:"rebuild_enabled,type:bit"`
}
