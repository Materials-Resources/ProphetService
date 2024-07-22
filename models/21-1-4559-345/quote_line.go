package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type QuoteLine struct {
	bun.BaseModel    `bun:"table:quote_line"`
	QuoteLineUid     int32     `bun:"quote_line_uid,type:int,pk"`                                    // What is the unique identifier for this quote line?
	OeLineUid        int32     `bun:"oe_line_uid,type:int"`                                          // Internal unique value to identify an order line.
	LineCompleteFlag string    `bun:"line_complete_flag,type:char(1),default:('N')"`                 // Identifies whether the line item has been marked as Complete by the user.  Should only contain values of Y & N.
	QtyConverted     float64   `bun:"qty_converted,type:decimal(19,9)"`                              // Accumulator/summation column that identifies the unit_quantity previously converted from the current quote to order(s).
	UomConverted     string    `bun:"uom_converted,type:varchar(8)"`                                 // The unit of measure associated with the qty_converted value.  From a business perspective -  should be the same as the unit_of_measure.
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`                // Indicates the date/time this record was created.
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`          // Indicates the date/time this record was last modified.
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"` // ID of the user who last maintained this record
}
