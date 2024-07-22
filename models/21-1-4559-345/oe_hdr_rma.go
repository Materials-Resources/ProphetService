package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type OeHdrRma struct {
	bun.BaseModel          `bun:"table:oe_hdr_rma"`
	OeHdrRmaUid            int32     `bun:"oe_hdr_rma_uid,type:int,pk"`                                   // Unique ID for this oe_hdr_rma record.
	OeHdrUid               int32     `bun:"oe_hdr_uid,type:int,unique"`                                   // FK to column oe_hdr.oe_hdr_uid.  Unique internal ID for the oe_hdr associated with this oe_hdr_rma record.
	ApplyCreditToInvoiceNo string    `bun:"apply_credit_to_invoice_no,type:char(1),default:('N')"`        // Indicates whether any associated credit memo
	DateCreated            time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy              string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified       time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy       string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	RetrievedByWms         string    `bun:"retrieved_by_wms,type:char(1),default:('N')"`                  // Determines whether or not this record has been exported.
}
