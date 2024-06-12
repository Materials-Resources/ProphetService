package model

import (
	"github.com/uptrace/bun"
	"time"
)

type PortalAssignment struct {
	bun.BaseModel       `bun:"table:portal_assignment"`
	PortalAssignmentUid int32     `bun:"portal_assignment_uid,type:int,pk"`
	PortalUid           int32     `bun:"portal_uid,type:int"`
	UserId              string    `bun:"user_id,type:varchar(30),nullzero"`
	RoleUid             int32     `bun:"role_uid,type:int,nullzero"`
	DateCreated         time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy           string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified    time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy    string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	RowStatusFlag       int32     `bun:"row_status_flag,type:int,default:((704))"`
	EnabledForVersionCd int32     `bun:"enabled_for_version_cd,type:int,default:((1423))"`
	RefreshTimer        int32     `bun:"refresh_timer,type:int,default:((0))"`
}
