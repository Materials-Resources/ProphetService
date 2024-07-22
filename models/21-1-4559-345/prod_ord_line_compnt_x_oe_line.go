package prophet

import (
	"github.com/uptrace/bun"
	"time"
)

type ProdOrdLineCompntXOeLine struct {
	bun.BaseModel             `bun:"table:prod_ord_line_compnt_x_oe_line"`
	ProdOrderLineComponentUid int32     `bun:"prod_order_line_component_uid,type:int,pk"`                 // Unique ID for this production order line component.
	OeLineUid                 int32     `bun:"oe_line_uid,type:int,pk"`                                   // Internal unique value to identify an order line.
	RowStatusFlag             int32     `bun:"row_status_flag,type:int"`                                  // Indicates current record status.
	DateCreated               time.Time `bun:"date_created,type:datetime"`                                // Indicates the date/time this record was created.
	DateLastModified          time.Time `bun:"date_last_modified,type:datetime"`                          // Indicates the date/time this record was last modified.
	LastMaintainedBy          string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"` // ID of the user who last maintained this record
}
