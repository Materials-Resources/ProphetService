package model

import (
	"github.com/uptrace/bun"
	"time"
)

type CarrierProviderType struct {
	bun.BaseModel           `bun:"table:carrier_provider_type"`
	CarrierProviderTypeUid  int32     `bun:"carrier_provider_type_uid,type:int,pk,identity"`
	CarrierProviderTypeId   int32     `bun:"carrier_provider_type_id,type:int"`
	CarrierProviderTypeDesc string    `bun:"carrier_provider_type_desc,type:varchar(255)"`
	DateCreated             time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy               string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
