package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type ItemConversion struct {
	bun.BaseModel          `bun:"table:item_conversion"`
	FromUom                string    `bun:"from_uom,type:varchar(8),pk"`                               // What unit of measure are we converting from?
	ToUom                  string    `bun:"to_uom,type:varchar(8),pk"`                                 // What is the unit of measure that we are converting to?
	Round                  string    `bun:"round,type:char(1)"`                                        // What kind of rounding should this conversion use? "U" = always round up -  "D" = always round down -  "N" = ? and "S" = ?
	Accumulate             string    `bun:"accumulate,type:char(1)"`                                   // Indicates whether to convert the Source UOM into a
	PrintOnPickTicket      string    `bun:"print_on_pick_ticket,type:char(1)"`                         // Should this conversion be printed on pick tickets?
	PrintOnOrderAck        string    `bun:"print_on_order_ack,type:char(1)"`                           // Should the conversion be printed on Order Acknowledgements?
	PrintOnInvoice         string    `bun:"print_on_invoice,type:char(1)"`                             // Deleted on or before 10/10/98 dstrait
	PrintOnPo              string    `bun:"print_on_po,type:char(1)"`                                  // Deleted on or before 10/10/98 dstrait
	ConvertAtOe            string    `bun:"convert_at_oe,type:char(1)"`                                // Display Conversion at Order Entry
	ConvertAtPo            string    `bun:"convert_at_po,type:char(1)"`                                // Display conversion at Purchase Order Entry
	DeleteFlag             string    `bun:"delete_flag,type:char(1)"`                                  // Indicates whether this record is logically deleted
	DateCreated            time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified       time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy       string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
	SortOrder              float64   `bun:"sort_order,type:decimal(19,0),pk"`                          // What display order should this item conversion appear in?
	InvMastUid             int32     `bun:"inv_mast_uid,type:int,pk"`                                  // Unique identifier for the item id.
	PrintOnProdOrder       string    `bun:"print_on_prod_order,type:char(1),default:('N')"`            // Should conversion be printed on production orders?
	ConvertAtTransferFlag  string    `bun:"convert_at_transfer_flag,type:char(1),default:('N')"`       // Flag indicating whether UOM conversion should happen on transfer
	MinOrderQty            *float64  `bun:"min_order_qty,type:decimal(19,9)"`                          // Custom - Minimum order quantity to be used when the conversion is applied in OE
	ConvertAtProdOrderFlag string    `bun:"convert_at_prod_order_flag,type:char(1),default:('N')"`     // Flag indicating whether UOM conversion should happen on production order entry
}
