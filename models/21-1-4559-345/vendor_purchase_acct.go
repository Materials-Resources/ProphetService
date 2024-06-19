package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type VendorPurchaseAcct struct {
	bun.BaseModel             `bun:"table:vendor_purchase_acct"`
	CompanyId                 string    `bun:"company_id,type:varchar(8),pk"`                                 // Unique code that identifies a company.
	VendorId                  float64   `bun:"vendor_id,type:decimal(18,0),pk"`                               // What is the vendor ID?
	PurchaseAccountNo         string    `bun:"purchase_account_no,type:varchar(80),pk"`                       // What is the vendors purchase account?
	DeleteFlag                string    `bun:"delete_flag,type:char(1)"`                                      // Indicates whether this record is logically deleted
	DateCreated               time.Time `bun:"date_created,type:datetime"`                                    // Indicates the date/time this record was created.
	DateLastModified          time.Time `bun:"date_last_modified,type:datetime"`                              // Indicates the date/time this record was last modified.
	LastMaintainedBy          string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"` // ID of the user who last maintained this record
	CoreAcctFlag              string    `bun:"core_acct_flag,type:char(1),nullzero"`                          // This custom column indicates the vendor purchase account is a core expense account
	FreightEstimatedAcctFlag  string    `bun:"freight_estimated_acct_flag,type:char(1),nullzero"`             // Custom column to indicate the account for estimated freight amount
	FreightDifferenceAcctFlag string    `bun:"freight_difference_acct_flag,type:char(1),nullzero"`            // Custom column to indicate the account for the difference between estimated freight amount and the actual freight amount.
	BranchId                  string    `bun:"branch_id,type:varchar(8),nullzero"`                            // Indicates the unique branch assoicated with the purchase account.
	AllocationPercent         float64   `bun:"allocation_percent,type:decimal(19,9),nullzero"`                // Indicates the percent of the invoice amount that will be allocated to the asociated branch.
	PurchaseDesc              string    `bun:"purchase_desc,type:varchar(255),nullzero"`                      // A default purchase description for doc-link's Smart Form.
	PurchaseCompanyId         string    `bun:"purchase_company_id,type:varchar(8),nullzero"`                  // (Custom) Indicates the unique company associated with the purchase account.
}
