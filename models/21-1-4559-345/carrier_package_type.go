package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type CarrierPackageType struct {
	bun.BaseModel          `bun:"table:carrier_package_type"`
	CarrierPackageTypeUid  int32     `bun:"carrier_package_type_uid,type:int,autoincrement,pk"`           // UID for table
	CarrierPackageTypeId   int32     `bun:"carrier_package_type_id,type:int"`                             // Name of carrier package type
	CarrierPackageTypeDesc string    `bun:"carrier_package_type_desc,type:varchar(255)"`                  // Description of carrier package type
	CarrierProviderTypeUid int32     `bun:"carrier_provider_type_uid,type:int"`                           // UID of carrier associated with package type
	DateCreated            time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy              string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified       time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy       string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
