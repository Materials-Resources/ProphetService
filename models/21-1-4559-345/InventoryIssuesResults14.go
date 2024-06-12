package model

import "github.com/uptrace/bun"

type Inventoryissuesresults14 struct {
	bun.BaseModel `bun:"table:InventoryIssuesResults14"`
	Run           int32   `bun:"run,type:int"`
	LocationId    float64 `bun:"location_id,type:decimal(19,0)"`
	ItemId        string  `bun:"item_id,type:varchar(40)"`
	InvMastUid    int32   `bun:"inv_mast_uid,type:int"`
	QtyOnHand     float64 `bun:"qty_on_hand,type:decimal(19,9),nullzero"`
}
