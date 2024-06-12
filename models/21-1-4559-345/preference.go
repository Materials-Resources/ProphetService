package model

import (
	"github.com/uptrace/bun"
	"time"
)

type Preference struct {
	bun.BaseModel    `bun:"table:preference"`
	PreferenceUid    int32     `bun:"preference_uid,type:int,pk,identity"`
	Name             string    `bun:"name,type:varchar(255)"`
	DefaultValue     string    `bun:"default_value,type:varchar"`
	DataTypeCd       int32     `bun:"data_type_cd,type:int,default:((1255))"`
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
