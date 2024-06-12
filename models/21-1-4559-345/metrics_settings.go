package model

import (
	"github.com/uptrace/bun"
	"time"
)

type MetricsSettings struct {
	bun.BaseModel    `bun:"table:metrics_settings"`
	SettingUid       int32     `bun:"setting_uid,type:int,pk,identity"`
	MajorKey         string    `bun:"major_key,type:nvarchar(50)"`
	MinorKey         string    `bun:"minor_key,type:nvarchar(50)"`
	Value            string    `bun:"value,type:nvarchar,nullzero"`
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy        string    `bun:"created_by,type:nvarchar(255),default:(suser_sname())"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:nvarchar(255),default:(suser_sname())"`
}
