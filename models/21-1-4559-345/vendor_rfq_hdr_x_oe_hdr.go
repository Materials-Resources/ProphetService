package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type VendorRfqHdrXOeHdr struct {
	bun.BaseModel         `bun:"table:vendor_rfq_hdr_x_oe_hdr"`
	VendorRfqHdrXOeHdrUid int32     `bun:"vendor_rfq_hdr_x_oe_hdr_uid,type:int,pk"`                      // Unique ID for each record in the table.
	OeHdrUid              int32     `bun:"oe_hdr_uid,type:int,unique"`                                   // Identifies unique Order record
	VendorRfqHdrUid       int32     `bun:"vendor_rfq_hdr_uid,type:int,unique"`                           // Identifies unique RFQ record
	DateCreated           time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy             string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified      time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy      string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
}
