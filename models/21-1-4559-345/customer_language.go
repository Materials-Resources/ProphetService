package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type CustomerLanguage struct {
	bun.BaseModel       `bun:"table:customer_language"`
	CustomerLanguageUid int32     `bun:"customer_language_uid,type:int,autoincrement,scanonly,pk"`     // Unique id for company_language record
	CompanyId           string    `bun:"company_id,type:varchar(8)"`                                   // Unique id for company code
	CustomerId          float64   `bun:"customer_id,type:decimal(19,0)"`                               // Unique id for customer.
	LanguageId          string    `bun:"language_id,type:varchar(8)"`                                  // A customer specified language for customer
	DateCreated         time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy           string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified    time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy    string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
