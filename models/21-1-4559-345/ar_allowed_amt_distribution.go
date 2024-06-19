package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type ArAllowedAmtDistribution struct {
	bun.BaseModel       `bun:"table:ar_allowed_amt_distribution"`
	ArAllowedAmtDistUid int32     `bun:"ar_allowed_amt_dist_uid,type:int,pk"`                          // Unique ID.
	ReceiptNumber       float64   `bun:"receipt_number,type:decimal(19,0)"`                            // Receipt number corresponding to payment at which time this allowed amt was taken.
	GlAllowedAcct       string    `bun:"gl_allowed_acct,type:varchar(32)"`                             // GL acct which allowed amt posted to.
	AllowedAmount       float64   `bun:"allowed_amount,type:decimal(19,2)"`                            // Allowed amount that was distributed to a particular gl acct.
	InvoiceNo           string    `bun:"invoice_no,type:varchar(10)"`                                  // Invoice for which this allowed amount was taken.
	DateCreated         time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy           string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified    time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy    string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
