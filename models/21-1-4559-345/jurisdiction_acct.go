package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type JurisdictionAcct struct {
	bun.BaseModel    `bun:"table:jurisdiction_acct"`
	CompanyNo        string    `bun:"company_no,type:varchar(8),pk"`                                 // Unique code that identifies a company.
	JurisdictionId   string    `bun:"jurisdiction_id,type:varchar(10),pk"`                           // What is the tax jurisdiction for this invoice line?
	GlAccountNo      string    `bun:"gl_account_no,type:varchar(32)"`                                // General Ledger Account number
	DeleteFlag       string    `bun:"delete_flag,type:char(1)"`                                      // Indicates whether this record is logically deleted
	DateCreated      time.Time `bun:"date_created,type:datetime"`                                    // Indicates the date/time this record was created.
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`                              // Indicates the date/time this record was last modified.
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"` // ID of the user who last maintained this record
}
