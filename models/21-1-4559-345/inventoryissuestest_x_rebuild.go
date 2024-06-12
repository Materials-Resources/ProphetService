package model

import (
	"github.com/uptrace/bun"
	"time"
)

type InventoryissuestestXRebuild struct {
	bun.BaseModel    `bun:"table:inventoryissuestest_x_rebuild"`
	Testuid          int32     `bun:"testuid,type:int"`
	Rebuilduid       int32     `bun:"rebuilduid,type:int"`
	SequenceNo       int32     `bun:"sequence_no,type:int,default:((1))"`
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	TotalTestUid     int32     `bun:"total_test_uid,type:int,nullzero"`
}
