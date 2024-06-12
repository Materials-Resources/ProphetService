package model

import (
	"github.com/uptrace/bun"
	"time"
)

type InvoiceHdrXTaxJurisManual struct {
	bun.BaseModel    `bun:"table:invoice_hdr_x_tax_juris_manual"`
	InvoiceNo        string    `bun:"invoice_no,type:varchar(10),pk"`
	JurisdictionId   string    `bun:"jurisdiction_id,type:varchar(10),pk"`
	Taxable          string    `bun:"taxable,type:char"`
	TaxPercentage    float64   `bun:"tax_percentage,type:decimal(19,8)"`
	DateCreated      time.Time `bun:"date_created,type:datetime"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
}
