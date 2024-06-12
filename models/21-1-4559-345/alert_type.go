package model

import (
	"github.com/uptrace/bun"
	"time"
)

type AlertType struct {
	bun.BaseModel    `bun:"table:alert_type"`
	AlertTypeUid     int32     `bun:"alert_type_uid,type:int,pk"`
	TypeCd           int32     `bun:"type_cd,type:int"`
	RowStatusFlag    int32     `bun:"row_status_flag,type:int"`
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	ModuleCd         int32     `bun:"module_cd,type:int"`
	ConfigurationId  int32     `bun:"configuration_id,type:int"`
	ViewName         string    `bun:"view_name,type:varchar(255)"`
	JobId            string    `bun:"job_id,type:varchar(255),nullzero"`
}
