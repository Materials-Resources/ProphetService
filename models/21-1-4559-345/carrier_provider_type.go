package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type CarrierProviderType struct {
	bun.BaseModel           `bun:"table:carrier_provider_type"`
	CarrierProviderTypeUid  int32     `bun:"carrier_provider_type_uid,type:int,autoincrement,pk"`          // Uniquely Indentifies a record in the table
	CarrierProviderTypeId   int32     `bun:"carrier_provider_type_id,type:int"`                            // Currently stores Agile specific code
	CarrierProviderTypeDesc string    `bun:"carrier_provider_type_desc,type:varchar(255)"`                 // Stores Provider Type description
	DateCreated             time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy               string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
