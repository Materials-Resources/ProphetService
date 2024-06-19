package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type CompanyLanguage struct {
	bun.BaseModel      `bun:"table:company_language"`
	CompanyLanguageUid int32     `bun:"company_language_uid,type:int,autoincrement,identity,pk"`      // Unique id for company_language record
	CompanyId          string    `bun:"company_id,type:varchar(8)"`                                   // Unique id for company code
	LanguageId         string    `bun:"language_id,type:varchar(8),nullzero"`                         // A default language for company
	DateCreated        time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy          string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified   time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy   string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
