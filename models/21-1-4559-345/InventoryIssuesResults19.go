package model

import "github.com/uptrace/bun"

type Inventoryissuesresults19 struct {
	bun.BaseModel         `bun:"table:InventoryIssuesResults19"`
	Run                   int32   `bun:"run,type:int"`
	LocationId            float64 `bun:"location_id,type:decimal(19,0)"`
	ItemId                string  `bun:"item_id,type:varchar(40)"`
	InvMastUid            int32   `bun:"inv_mast_uid,type:int"`
	InvLocQtyAllocated    float64 `bun:"inv_loc_qty_allocated,type:decimal(19,9),nullzero"`
	SumInvBinQtyAllocated float64 `bun:"sum_inv_bin_qty_allocated,type:decimal(38,9),nullzero"`
}
