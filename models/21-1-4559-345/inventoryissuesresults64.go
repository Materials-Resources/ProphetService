package prophet

import "github.com/uptrace/bun"

type Inventoryissuesresults64 struct {
	bun.BaseModel            `bun:"table:inventoryissuesresults64"`
	Run                      int32    `bun:"run,type:int"`                                    // Run number
	LocationId               int32    `bun:"location_id,type:int"`                            // Location ID
	ItemId                   string   `bun:"item_id,type:varchar(40)"`                        // Item ID
	InvMastUid               int32    `bun:"inv_mast_uid,type:int"`                           // inv_mast_uid
	SumTranLinesQtyAllocated *float64 `bun:"sum_tran_lines_qty_allocated,type:decimal(38,9)"` // Total Qty Allocated (Transactions)
	TableInvLocQtyAllocated  *float64 `bun:"table_inv_loc_qty_allocated,type:decimal(38,9)"`  // inv_loc.qty_allocated
}
