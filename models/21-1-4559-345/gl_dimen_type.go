package model

import (
	"github.com/uptrace/bun"
	"time"
)

type GlDimenType struct {
	bun.BaseModel    `bun:"table:gl_dimen_type"`
	GlDimenTypeUid   int32     `bun:"gl_dimen_type_uid,type:int,pk,identity"`
	GlDimenTypeId    string    `bun:"gl_dimen_type_id,type:varchar(255)"`
	GlDimenTypeDesc  string    `bun:"gl_dimen_type_desc,type:varchar(255)"`
	RecordTypeCd     int32     `bun:"record_type_cd,type:int,default:((932))"`
	DeleteFlag       string    `bun:"delete_flag,type:char,default:('N')"`
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	UseValues        string    `bun:"use_values,type:char,default:('N')"`
}
