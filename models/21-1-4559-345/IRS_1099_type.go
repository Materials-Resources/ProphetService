package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type Irs1099Type struct {
	bun.BaseModel         `bun:"table:IRS_1099_type"`
	Irs1099TypeUid        int32     `bun:"IRS_1099_type_uid,type:int,pk"`                                // Unique ID of this 1099 type
	DisplayValue          string    `bun:"display_value,type:varchar(255)"`                              // The name of this 1099 type
	Threshold             float64   `bun:"threshold,type:decimal(19,2),default:(0)"`                     // Threshold for this 1099 type
	RowStatusFlag         int32     `bun:"row_status_flag,type:int,default:(704)"`                       // Status of the row.
	DateCreated           time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy             string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified      time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy      string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	AllowIncorporatedFlag string    `bun:"allow_incorporated_flag,type:char(1),default:('N')"`           // Whether or not this 1099 type may be used by an incorporated vendor.
}
