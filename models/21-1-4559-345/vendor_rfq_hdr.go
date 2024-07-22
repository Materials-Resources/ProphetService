package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type VendorRfqHdr struct {
	bun.BaseModel    `bun:"table:vendor_rfq_hdr"`
	VendorRfqHdrUid  int32      `bun:"vendor_rfq_hdr_uid,type:int,pk"`                               // Unique ID for each record in the table.
	PoHdrUid         int32      `bun:"po_hdr_uid,type:int,unique"`                                   // Identifies unique PO record
	ControlNo        int32      `bun:"control_no,type:int"`                                          // Groups RFQ, PO and Sales Orders
	ExpirationDate   time.Time  `bun:"expiration_date,type:datetime"`                                // Date RFQ will expire
	DateCreated      time.Time  `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy        string     `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified time.Time  `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy string     `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	Date840Sent      *time.Time `bun:"date_840_sent,type:datetime"`                                  // Last Send of EDI 840 PO RFQ
	TypeCd           *int32     `bun:"type_cd,type:int"`                                             // Code to identify type of Vendor RFQ.
	ResponseDate     *time.Time `bun:"response_date,type:datetime,default:(null)"`                   // Date received 843 response to 840
}
