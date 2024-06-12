package model

import (
	"github.com/uptrace/bun"
	"time"
)

type PricingTemplate struct {
	bun.BaseModel       `bun:"table:pricing_template"`
	PricingTemplateUid  int32     `bun:"pricing_template_uid,type:int,pk"`
	PricingTemplateId   string    `bun:"pricing_template_id,type:varchar(255)"`
	PricingTemplateDesc string    `bun:"pricing_template_desc,type:varchar(255)"`
	KeyFieldCd          int32     `bun:"key_field_cd,type:int"`
	RowStatusFlag       int32     `bun:"row_status_flag,type:int"`
	DateCreated         time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy           string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified    time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy    string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	SecondaryKeyFieldCd int32     `bun:"secondary_key_field_cd,type:int,nullzero"`
}
