package model

import (
	"github.com/uptrace/bun"
	"time"
)

type BusinessObjectKeyFields struct {
	bun.BaseModel              `bun:"table:business_object_key_fields"`
	BusinessObjectKeyFieldsUid int32     `bun:"business_object_key_fields_uid,type:int,pk,identity"`
	BusinessObjectName         string    `bun:"business_object_name,type:varchar(100)"`
	KeyFieldNames              string    `bun:"key_field_names,type:varchar(255)"`
	DateCreated                time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                  string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified           time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy           string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
