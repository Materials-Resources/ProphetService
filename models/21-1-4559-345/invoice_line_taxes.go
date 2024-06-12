package model

import (
	"github.com/uptrace/bun"
	"time"
)

type InvoiceLineTaxes struct {
	bun.BaseModel       `bun:"table:invoice_line_taxes"`
	InvoiceNo           string    `bun:"invoice_no,type:varchar(10),pk"`
	ItemId              string    `bun:"item_id,type:varchar(40)"`
	JurisdictionId      string    `bun:"jurisdiction_id,type:varchar(10),pk"`
	TaxPercentage       float64   `bun:"tax_percentage,type:decimal(19,8)"`
	DateCreated         time.Time `bun:"date_created,type:datetime"`
	DateLastModified    time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy    string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"`
	LineNo              float64   `bun:"line_no,type:decimal(19,0),pk"`
	Taxable             string    `bun:"taxable,type:char"`
	TaxCharged          float64   `bun:"tax_charged,type:decimal(19,4),default:(0)"`
	TotalSalesAmt       float64   `bun:"total_sales_amt,type:decimal(19,4),default:(0)"`
	TaxableSalesAmt     float64   `bun:"taxable_sales_amt,type:decimal(19,4),default:(0)"`
	InvoiceLineTaxesUid int32     `bun:"invoice_line_taxes_uid,type:int,identity"`
}
