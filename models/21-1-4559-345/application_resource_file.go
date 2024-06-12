package model

import (
	"github.com/uptrace/bun"
	"time"
)

type ApplicationResourceFile struct {
	bun.BaseModel              `bun:"table:application_resource_file"`
	ApplicationResourceFileUid int32     `bun:"application_resource_file_uid,type:int,pk,identity"`
	FileName                   string    `bun:"file_name,type:varchar(255)"`
	FileType                   string    `bun:"file_type,type:varchar(255),default:('IMAGE')"`
	DateCreated                time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                  string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified           time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy           string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
