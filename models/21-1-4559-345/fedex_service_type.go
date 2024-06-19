package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type FedexServiceType struct {
	bun.BaseModel        `bun:"table:fedex_service_type"`
	FedexServiceTypeUid  int32     `bun:"fedex_service_type_uid,type:int,pk"`                           // Surrogate key for the table
	FedexServiceTypeId   string    `bun:"fedex_service_type_id,type:varchar(255),unique"`               // Logical key for the table
	FedexServiceTypeDesc string    `bun:"fedex_service_type_desc,type:varchar(255)"`                    // description
	RowStatusFlag        int32     `bun:"row_status_flag,type:int,default:(704)"`                       // indicates status (active, deleted, etc.)
	DateCreated          time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy            string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified     time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy     string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	FedexExpressFlag     string    `bun:"fedex_express_flag,type:char(1),default:('N')"`                // Each service type needs to be a Fedex ground or express transaction
}
