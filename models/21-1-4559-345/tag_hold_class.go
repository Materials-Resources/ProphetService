package model

import (
	"github.com/uptrace/bun"
	"time"
)

type TagHoldClass struct {
	bun.BaseModel    `bun:"table:tag_hold_class"`
	TagHoldClassUid  int32     `bun:"tag_hold_class_uid,type:int,pk"`
	TagHoldClassId   string    `bun:"tag_hold_class_id,type:varchar(255)"`
	TagHoldClassDesc string    `bun:"tag_hold_class_desc,type:varchar(255)"`
	RowStatusFlag    int32     `bun:"row_status_flag,type:int"`
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
