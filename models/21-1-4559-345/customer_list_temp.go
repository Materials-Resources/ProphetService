package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type CustomerListTemp struct {
	bun.BaseModel       `bun:"table:customer_list_temp"`
	CustomerListTempUid int32     `bun:"customer_list_temp_uid,type:int,autoincrement,identity"`       // Unique identifier
	RunNumber           int32     `bun:"run_number,type:int"`                                          // Creates a batch of customers
	CompanyId           string    `bun:"company_id,type:varchar(8)"`                                   // Company ID associated with the batch
	CustomerId          float64   `bun:"customer_id,type:decimal(19,0)"`                               // Customer in the batch
	DateCreated         time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy           string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified    time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy    string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
