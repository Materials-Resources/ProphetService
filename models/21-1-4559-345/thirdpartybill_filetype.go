package model

import (
	"github.com/uptrace/bun"
	"time"
)

type ThirdpartybillFiletype struct {
	bun.BaseModel             `bun:"table:thirdpartybill_filetype"`
	ThirdpartybillFiletypeUid int32     `bun:"thirdpartybill_filetype_uid,type:int,pk,identity"`
	FiletypePrefix            string    `bun:"filetype_prefix,type:varchar(255)"`
	Description               string    `bun:"description,type:varchar(255)"`
	RowStatusFlag             int32     `bun:"row_status_flag,type:int"`
	DateCreated               time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                 string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified          time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy          string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
