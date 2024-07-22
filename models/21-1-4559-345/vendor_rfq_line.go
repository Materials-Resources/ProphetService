package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type VendorRfqLine struct {
	bun.BaseModel           `bun:"table:vendor_rfq_line"`
	VendorRfqLineUid        int32     `bun:"vendor_rfq_line_uid,type:int,pk"`                              // Unique ID for each record in the table.
	PoLineUid               int32     `bun:"po_line_uid,type:int,unique"`                                  // Identifies unique PO line record
	QtyConverted            float64   `bun:"qty_converted,type:decimal(19,9),default:(0)"`                 // Keeps running total of qty converted to a PO
	DateCreated             time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy               string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	DeliveryInfo            *string   `bun:"delivery_info,type:varchar(255)"`                              // Stores RFQ delivery information per line
	Responded               int32     `bun:"responded,type:int,default:(1104)"`                            // State of Vendor RFQ response
	RfqanalysisCompleteFlag *string   `bun:"rfqanalysis_complete_flag,type:char(1)"`                       // Determines if this RFQ should be excluded from future RFQ to Sales Quote analysis.
}
