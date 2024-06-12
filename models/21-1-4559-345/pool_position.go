package model

import (
	"github.com/uptrace/bun"
	"time"
)

type PoolPosition struct {
	bun.BaseModel    `bun:"table:pool_position"`
	PoolPositionUid  int32     `bun:"pool_position_uid,type:int,pk,identity"`
	PoolPositionId   string    `bun:"pool_position_id,type:varchar(20)"`
	Description      string    `bun:"description,type:varchar(255)"`
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
