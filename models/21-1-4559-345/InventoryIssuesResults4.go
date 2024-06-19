package gen

import "github.com/uptrace/bun"

type Inventoryissuesresults4 struct {
	bun.BaseModel      `bun:"table:InventoryIssuesResults4"`
	Run                int32   `bun:"run,type:int"`
	LocationId         float64 `bun:"location_id,type:decimal(19,0)"`
	ItemId             string  `bun:"item_id,type:varchar(40)"`
	InvMastUid         int32   `bun:"inv_mast_uid,type:int"`
	InvLocQtyOnHand    float64 `bun:"inv_loc_qty_on_hand,type:decimal(19,9),nullzero"`
	SumInvBinQtyOnHand float64 `bun:"sum_inv_bin_qty_on_hand,type:decimal(38,9),nullzero"`
}
