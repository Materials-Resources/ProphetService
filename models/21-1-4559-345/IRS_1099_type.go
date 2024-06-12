package model

import (
	"github.com/uptrace/bun"
	"time"
)

type Irs1099Type struct {
	bun.BaseModel         `bun:"table:IRS_1099_type"`
	Irs1099TypeUid        int32     `bun:"IRS_1099_type_uid,type:int,pk"`
	DisplayValue          string    `bun:"display_value,type:varchar(255)"`
	Threshold             float64   `bun:"threshold,type:decimal(19,2),default:(0)"`
	RowStatusFlag         int32     `bun:"row_status_flag,type:int,default:(704)"`
	DateCreated           time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy             string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified      time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy      string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	AllowIncorporatedFlag string    `bun:"allow_incorporated_flag,type:char,default:('N')"`
}
