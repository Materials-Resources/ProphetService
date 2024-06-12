package model

import (
	"github.com/uptrace/bun"
	"time"
)

type SystemSetting struct {
	bun.BaseModel    `bun:"table:system_setting"`
	SystemSettingUid int32     `bun:"system_setting_uid,type:int,pk"`
	ConfigurationId  int32     `bun:"configuration_id,type:int"`
	ModuleCd         int32     `bun:"module_cd,type:int"`
	Name             string    `bun:"name,type:varchar(128)"`
	Value            string    `bun:"value,type:varchar(255)"`
	DataTypeCd       int32     `bun:"data_type_cd,type:int"`
	DataTypeLength   int32     `bun:"data_type_length,type:int"`
	DataTypeScale    int32     `bun:"data_type_scale,type:int"`
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
}
