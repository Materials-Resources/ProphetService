package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type ProdOrdLinePo struct {
	bun.BaseModel    `bun:"table:prod_ord_line_po"`
	ProdOrderNumber  float64   `bun:"prod_order_number,type:decimal(19,0),pk"`       // Production Order associated with this prod_ord_line_po record
	LineNumber       float64   `bun:"line_number,type:decimal(19,0),pk"`             // Production Order line number associated with this prod_ord_line_po record
	ComponentNumber  float64   `bun:"component_number,type:decimal(19,0),pk"`        // Production Order component number associated with this prod_ord_line_po record
	PoNumber         float64   `bun:"po_number,type:decimal(19,0),pk"`               // Purchase Order associated with this prod_ord_line_po record
	PoLineNumber     float64   `bun:"po_line_number,type:decimal(19,0),pk"`          // Purchase Order line associated with this prod_ord_line_po record
	Quantity         float64   `bun:"quantity,type:decimal(19,9)"`                   // Quantity for this component tied to the purchase order
	DateCreated      time.Time `bun:"date_created,type:datetime"`                    // Indicates the date/time this record was created.
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`              // Indicates the date/time this record was last modified.
	RowStatusFlag    int16     `bun:"row_status_flag,type:smallint,default:(1)"`     // Indicates current record status.
	ConnectionType   string    `bun:"connection_type,type:char(1),default:('P'),pk"` // Indicates the type of transaction connected to the Production Order.
}
