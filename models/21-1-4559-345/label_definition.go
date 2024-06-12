package model

import (
	"github.com/uptrace/bun"
	"time"
)

type LabelDefinition struct {
	bun.BaseModel       `bun:"table:label_definition"`
	LabelDefinitionUid  int32     `bun:"label_definition_uid,type:int,pk"`
	CompanyId           string    `bun:"company_id,type:varchar(8)"`
	LabelDefinitionId   string    `bun:"label_definition_id,type:varchar(255)"`
	LabelDefinitionDesc string    `bun:"label_definition_desc,type:varchar(255),nullzero"`
	LabelTypeCd         int32     `bun:"label_type_cd,type:int"`
	RowStatusFlag       int32     `bun:"row_status_flag,type:int"`
	DateCreated         time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy           string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified    time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy    string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
