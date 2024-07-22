package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type OeHdrTax struct {
	bun.BaseModel    `bun:"table:oe_hdr_tax"`
	OrderNo          string    `bun:"order_no,type:varchar(8),pk"`                                   // What order does this invoice belong to?
	JurisdictionId   string    `bun:"jurisdiction_id,type:varchar(10),pk"`                           // What is the unique identifier for this jurisdiction?
	Taxable          string    `bun:"taxable,type:char(1)"`                                          // Is this ship to jurisdiction taxable?
	DateCreated      time.Time `bun:"date_created,type:datetime"`                                    // Indicates the date/time this record was created.
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`                              // Indicates the date/time this record was last modified.
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"` // ID of the user who last maintained this record
	RmaLinkedTax     *string   `bun:"rma_linked_tax,type:char(1)"`                                   // Indicates whether a record was added as a result of linking an RMA line to an order
}
