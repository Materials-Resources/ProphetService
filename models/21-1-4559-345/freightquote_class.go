package model

import (
	"github.com/uptrace/bun"
	"time"
)

type FreightquoteClass struct {
	bun.BaseModel        `bun:"table:freightquote_class"`
	FreightquoteClassUid int32     `bun:"freightquote_class_uid,type:int,pk,identity"`
	FreightquoteClassId  float64   `bun:"freightquote_class_id,type:decimal(19,9)"`
	RowStatusFlag        int32     `bun:"row_status_flag,type:int"`
	DateCreated          time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy            string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified     time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy     string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
