package model

import (
	"github.com/uptrace/bun"
	"time"
)

type Portal struct {
	bun.BaseModel    `bun:"table:portal"`
	PortalUid        int32     `bun:"portal_uid,type:int,pk"`
	PortalName       string    `bun:"portal_name,type:varchar(255)"`
	PortalLayout     int32     `bun:"portal_layout,type:int"`
	PortalDesc       string    `bun:"portal_desc,type:varchar(255),nullzero"`
	RowStatusFlag    int32     `bun:"row_status_flag,type:int,default:((704))"`
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	RefreshTimer     int32     `bun:"refresh_timer,type:int,default:((0))"`
}
