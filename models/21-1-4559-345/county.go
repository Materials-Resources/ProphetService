package model

import (
	"github.com/uptrace/bun"
	"time"
)

type County struct {
	bun.BaseModel    `bun:"table:county"`
	CountyUid        int32     `bun:"county_uid,type:int,pk"`
	StateUid         int32     `bun:"state_uid,type:int"`
	CountyName       string    `bun:"county_name,type:varchar(255)"`
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
