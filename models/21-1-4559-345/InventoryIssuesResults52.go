package prophet

import "github.com/uptrace/bun"

type Inventoryissuesresults52 struct {
	bun.BaseModel    `bun:"table:InventoryIssuesResults52"`
	Run              int32    `bun:"run,type:int"`
	LocationId       float64  `bun:"location_id,type:decimal(19,0)"`
	ItemId           string   `bun:"item_id,type:varchar(40)"`
	InvMastUid       int32    `bun:"inv_mast_uid,type:int"`
	OrderNo          string   `bun:"order_no,type:varchar(8)"`
	LineNo           float64  `bun:"line_no,type:decimal(19,0)"`
	QtyOnPickTickets *float64 `bun:"qty_on_pick_tickets,type:decimal(19,9)"`
}
