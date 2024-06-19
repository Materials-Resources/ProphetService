package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type FedexSmartpostHub struct {
	bun.BaseModel         `bun:"table:fedex_smartpost_hub"`
	FedexSmartpostHubUid  int32     `bun:"fedex_smartpost_hub_uid,type:int,autoincrement,identity,pk"`   // Surrogate key for the table
	FedexSmartpostHubId   int32     `bun:"fedex_smartpost_hub_id,type:int,unique"`                       // Logical key for the table
	FedexSmartpostHubDesc string    `bun:"fedex_smartpost_hub_desc,type:varchar(255)"`                   // Description
	DateCreated           time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy             string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified      time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy      string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
