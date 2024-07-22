package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type County struct {
	bun.BaseModel    `bun:"table:county"`
	CountyUid        int32     `bun:"county_uid,type:int,pk"`                                       // Unique ID for county
	StateUid         int32     `bun:"state_uid,type:int,unique"`                                    // Unique ID for state
	CountyName       string    `bun:"county_name,type:varchar(255),unique"`                         // County name
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
