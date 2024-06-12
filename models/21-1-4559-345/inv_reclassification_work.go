package model

import (
	"github.com/uptrace/bun"
	"time"
)

type InvReclassificationWork struct {
	bun.BaseModel              `bun:"table:inv_reclassification_work"`
	InvReclassificationWorkUid int32     `bun:"inv_reclassification_work_uid,type:int,pk"`
	LocationId                 float64   `bun:"location_id,type:decimal(19,0)"`
	TotalItems                 int32     `bun:"total_items,type:int,default:(0)"`
	TotalValue                 float64   `bun:"total_value,type:decimal(19,2),default:(0)"`
	DateCreated                time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                  string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified           time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy           string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
