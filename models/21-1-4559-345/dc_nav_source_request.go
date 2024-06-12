package model

import (
	"github.com/uptrace/bun"
	"time"
)

type DcNavSourceRequest struct {
	bun.BaseModel         `bun:"table:dc_nav_source_request"`
	DcNavSourceRequestUid int32     `bun:"dc_nav_source_request_uid,type:int,pk,identity"`
	SourceWindow          string    `bun:"source_window,type:varchar(255)"`
	SourceDatawindow      string    `bun:"source_datawindow,type:varchar(255),nullzero"`
	SourceField           string    `bun:"source_field,type:varchar(255)"`
	DateCreated           time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy             string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified      time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy      string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	NavigationType        int32     `bun:"navigation_type,type:int,default:((3590))"`
}
