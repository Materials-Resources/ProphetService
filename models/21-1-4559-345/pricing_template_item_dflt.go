package model

import (
	"github.com/uptrace/bun"
	"time"
)

type PricingTemplateItemDflt struct {
	bun.BaseModel              `bun:"table:pricing_template_item_dflt"`
	PricingTemplateItemDfltUid int32     `bun:"pricing_template_item_dflt_uid,type:int,pk,identity"`
	PricingTemplateKeyFieldUid int32     `bun:"pricing_template_key_field_uid,type:int"`
	ColumnId                   float64   `bun:"column_id,type:decimal(19,0)"`
	DefaultValue               string    `bun:"default_value,type:varchar(255)"`
	DateCreated                time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                  string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified           time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy           string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
