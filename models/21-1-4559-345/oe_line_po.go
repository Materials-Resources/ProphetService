package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type OeLinePo struct {
	bun.BaseModel           `bun:"table:oe_line_po"`
	OrderNumber             string    `bun:"order_number,type:varchar(8),pk"`                           // Order number
	LineNumber              float64   `bun:"line_number,type:decimal(19,0),pk"`                         // The line number on the order.
	PoNo                    float64   `bun:"po_no,type:decimal(19,0),pk"`                               // Purchase order linked to order
	DateCreated             time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified        time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy        string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	QuantityOnPo            *float64  `bun:"quantity_on_po,type:decimal(19,9)"`                         // Maintains the portion of the actual po or transfer qty that is considered linked to a sales order.
	PoLineNumber            float64   `bun:"po_line_number,type:decimal(19,0),pk"`                      // Purchase order line linked to order line.
	DeleteFlag              string    `bun:"delete_flag,type:char(1)"`                                  // Indicates whether this record is logically deleted
	CancelFlag              *string   `bun:"cancel_flag,type:char(1)"`                                  // When Y this link has been canceled.
	Completed               string    `bun:"completed,type:char(1)"`                                    // When Y the purchase order has been received and the sales order allocated.
	ConnectionType          string    `bun:"connection_type,type:char(1),pk"`                           // Indicates whether the record is between an order and po or transfer and po.
	OriginalOeLineFreightIn *float64  `bun:"original_oe_line_freight_in,type:decimal(19,9)"`            // Custom column to save the original oe line freight in amount to oe_line_po
}
