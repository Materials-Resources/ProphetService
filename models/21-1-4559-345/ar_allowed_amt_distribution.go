package model

import (
	"github.com/uptrace/bun"
	"time"
)

type ArAllowedAmtDistribution struct {
	bun.BaseModel       `bun:"table:ar_allowed_amt_distribution"`
	ArAllowedAmtDistUid int32     `bun:"ar_allowed_amt_dist_uid,type:int,pk"`
	ReceiptNumber       float64   `bun:"receipt_number,type:decimal(19,0)"`
	GlAllowedAcct       string    `bun:"gl_allowed_acct,type:varchar(32)"`
	AllowedAmount       float64   `bun:"allowed_amount,type:decimal(19,2)"`
	InvoiceNo           string    `bun:"invoice_no,type:varchar(10)"`
	DateCreated         time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy           string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified    time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy    string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
