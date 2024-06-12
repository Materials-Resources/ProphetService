package model

import (
	"github.com/uptrace/bun"
	"time"
)

type ScheduledJobFeatureType struct {
	bun.BaseModel               `bun:"table:scheduled_job_feature_type"`
	ScheduledJobFeatureTypeUid  int32     `bun:"scheduled_job_feature_type_uid,type:int,pk,identity"`
	ScheduledJobFeatureTypeId   string    `bun:"scheduled_job_feature_type_id,type:varchar(255)"`
	ScheduledJobFeatureTypeDesc string    `bun:"scheduled_job_feature_type_desc,type:varchar(255)"`
	DeleteFlag                  string    `bun:"delete_flag,type:char,default:('N')"`
	DateCreated                 time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                   string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified            time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy            string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
