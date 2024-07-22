package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type VendorDefaults struct {
	bun.BaseModel    `bun:"table:vendor_defaults"`
	CompanyId        string    `bun:"company_id,type:varchar(8),pk"`                             // Unique code that identifies a company.
	ApAccountNumber  string    `bun:"ap_account_number,type:varchar(32)"`                        // What is the default AP account for vendors of this company?
	TermsId          string    `bun:"terms_id,type:varchar(2)"`                                  // What is the default Terms ID for vendors of this company?
	EdiOrPaper       string    `bun:"edi_or_paper,type:char(1)"`                                 // Does this customer use EDI or paper?
	CurrencyId       float64   `bun:"currency_id,type:decimal(19,0)"`                            // What is the default currency type for vendors of this company?
	DeleteFlag       string    `bun:"delete_flag,type:char(1)"`                                  // Indicates whether this record is logically deleted
	DateCreated      time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
}
