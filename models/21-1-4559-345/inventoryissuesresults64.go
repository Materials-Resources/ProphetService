package model

import "github.com/uptrace/bun"

type Inventoryissuesresults64 struct {
	bun.BaseModel            `bun:"table:inventoryissuesresults64"`
	Run                      int32   `bun:"run,type:int"`
	LocationId               int32   `bun:"location_id,type:int"`
	ItemId                   string  `bun:"item_id,type:varchar(40)"`
	InvMastUid               int32   `bun:"inv_mast_uid,type:int"`
	SumTranLinesQtyAllocated float64 `bun:"sum_tran_lines_qty_allocated,type:decimal(38,9),nullzero"`
	TableInvLocQtyAllocated  float64 `bun:"table_inv_loc_qty_allocated,type:decimal(38,9),nullzero"`
}
