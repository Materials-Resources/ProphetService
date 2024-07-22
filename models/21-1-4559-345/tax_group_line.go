package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type TaxGroupLine struct {
	bun.BaseModel    `bun:"table:tax_group_line"`
	TaxGroupId       string    `bun:"tax_group_id,type:varchar(10),pk"`                              // Indicates the tax group identification.
	CompanyId        string    `bun:"company_id,type:varchar(8),pk"`                                 // Unique code that identifies a company.
	JurisdictionId   string    `bun:"jurisdiction_id,type:varchar(10),pk"`                           // What is the unique identifier for this tax jurisdiction?
	DeleteFlag       string    `bun:"delete_flag,type:char(1)"`                                      // Indicates whether this record is logically deleted
	DateCreated      time.Time `bun:"date_created,type:datetime"`                                    // Indicates the date/time this record was created.
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`                              // Indicates the date/time this record was last modified.
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"` // ID of the user who last maintained this record
	Taxable          *string   `bun:"taxable,type:char(1)"`                                          // Is this line item taxable?
}
