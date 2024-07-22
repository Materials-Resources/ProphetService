package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type PoolShape struct {
	bun.BaseModel    `bun:"table:pool_shape"`
	PoolShapeUid     int32     `bun:"pool_shape_uid,type:int,autoincrement,identity,pk"`            // Unique ID for this table
	PoolShapeId      string    `bun:"pool_shape_id,type:varchar(20)"`                               // User defined code for this pool shape
	Description      string    `bun:"description,type:varchar(255)"`                                // Description for this pool shape
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
