package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type AiaElement struct {
	bun.BaseModel    `bun:"table:aia_element"`
	AiaElementUid    int32     `bun:"aia_element_uid,type:int,autoincrement,identity,pk"`           // Unique ID for this AIA element
	Description      string    `bun:"description,type:varchar(255)"`                                // The description of this AIA element
	PointValue       float64   `bun:"point_value,type:decimal(19,9),default:((0))"`                 // The point value of this AIA element
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
