package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type CarrierServiceType struct {
	bun.BaseModel          `bun:"table:carrier_service_type"`
	CarrierServiceTypeUid  int32     `bun:"carrier_service_type_uid,type:int,autoincrement,identity,pk"`  // UID for this table.
	CarrierServiceTypeId   int32     `bun:"carrier_service_type_id,type:int"`                             // Predefined value for this service type/carrier.
	CarrierServiceTypeDesc string    `bun:"carrier_service_type_desc,type:varchar(255)"`                  // Predefined desc for this service type/carrier value.
	CarrierProviderTypeUid int32     `bun:"carrier_provider_type_uid,type:int"`                           // UID of carrier for which this service type applies.
	DateCreated            time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy              string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified       time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy       string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
