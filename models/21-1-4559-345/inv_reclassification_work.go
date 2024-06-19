package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type InvReclassificationWork struct {
	bun.BaseModel              `bun:"table:inv_reclassification_work"`
	InvReclassificationWorkUid int32     `bun:"inv_reclassification_work_uid,type:int,pk"`                    // Uniquely Identifies each record in the table.
	LocationId                 float64   `bun:"location_id,type:decimal(19,0)"`                               // Location being ranked or reclassified
	TotalItems                 int32     `bun:"total_items,type:int,default:(0)"`                             // Total items ranked or reclassified at the location
	TotalValue                 float64   `bun:"total_value,type:decimal(19,2),default:(0)"`                   // Total value for the ranking / reclassification
	DateCreated                time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy                  string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified           time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy           string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
