package model

import (
	"github.com/uptrace/bun"
	"time"
)

type InvMastLinks struct {
	bun.BaseModel    `bun:"table:inv_mast_links"`
	InvMastLinksUid  int32     `bun:"inv_mast_links_uid,type:int,pk"`
	InvMastUid       int32     `bun:"inv_mast_uid,type:int"`
	LinkName         string    `bun:"link_name,type:varchar(30)"`
	LinkPath         string    `bun:"link_path,type:varchar(255)"`
	LinkArea         int32     `bun:"link_area,type:int"`
	RowStatusFlag    int32     `bun:"row_status_flag,type:int,default:(704)"`
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastModifiedBy   string    `bun:"last_modified_by,type:varchar(30),default:(user_name())"`
}
