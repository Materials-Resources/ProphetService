package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type InvoiceLineTaxes struct {
	bun.BaseModel       `bun:"table:invoice_line_taxes"`
	InvoiceNo           string    `bun:"invoice_no,type:varchar(10),pk"`                                // Invoice being paid
	ItemId              string    `bun:"item_id,type:varchar(40)"`                                      // The Item_id on the invoice_line record that this record is associated to
	JurisdictionId      string    `bun:"jurisdiction_id,type:varchar(10),pk"`                           // What is the unique identifier for this tax jurisdiction?
	TaxPercentage       float64   `bun:"tax_percentage,type:decimal(19,8)"`                             // Tax percentage of the jurisdiction at the time the record is created
	DateCreated         time.Time `bun:"date_created,type:datetime"`                                    // Indicates the date/time this record was created.
	DateLastModified    time.Time `bun:"date_last_modified,type:datetime"`                              // Indicates the date/time this record was last modified.
	LastMaintainedBy    string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"` // ID of the user who last maintained this record
	LineNo              float64   `bun:"line_no,type:decimal(19,0),pk"`                                 // The line number on the invoice_line record that this record is associated to
	Taxable             string    `bun:"taxable,type:char(1)"`                                          // Indicates if the jurisdiction was charged tax
	TaxCharged          float64   `bun:"tax_charged,type:decimal(19,4),default:(0)"`                    // What was the tax charged for this line -  for this jurisdiction?
	TotalSalesAmt       float64   `bun:"total_sales_amt,type:decimal(19,4),default:(0)"`                // The total cost of the items on an order.
	TaxableSalesAmt     float64   `bun:"taxable_sales_amt,type:decimal(19,4),default:(0)"`              // This column will hold the taxable sales for Tax Adjustments only.  Will be 0 for regular tax records.
	InvoiceLineTaxesUid int32     `bun:"invoice_line_taxes_uid,type:int,autoincrement,unique"`          // Unique identifier for a record.
}
