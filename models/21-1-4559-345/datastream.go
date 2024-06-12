package model

import (
	"github.com/uptrace/bun"
	"time"
)

type Datastream struct {
	bun.BaseModel    `bun:"table:datastream"`
	DatastreamUid    int32     `bun:"datastream_uid,type:int,pk,identity"`
	DatastreamName   string    `bun:"datastream_name,type:varchar(255)"`
	DatastreamDesc   string    `bun:"datastream_desc,type:varchar(255),nullzero"`
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
