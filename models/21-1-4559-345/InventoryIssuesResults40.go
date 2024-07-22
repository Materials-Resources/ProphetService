package prophet

import "github.com/uptrace/bun"

type Inventoryissuesresults40 struct {
	bun.BaseModel                      `bun:"table:InventoryIssuesResults40"`
	Run                                int32    `bun:"run,type:int"`
	LocationId                         float64  `bun:"location_id,type:decimal(19,0)"`
	ItemId                             string   `bun:"item_id,type:varchar(40)"`
	InvMastUid                         int32    `bun:"inv_mast_uid,type:int"`
	Bin                                string   `bun:"bin,type:varchar(10)"`
	SumTransactionLinesBinQtyAllocated *float64 `bun:"sum_transaction_lines_bin_qty_allocated,type:decimal(38,9)"`
	TableInvBinQtyAllocated            float64  `bun:"table_inv_bin_qty_allocated,type:decimal(19,9)"`
	BinId                              string   `bun:"bin_id,type:varchar(10)"`
	BinType                            *string  `bun:"bin_type,type:varchar(255)"`
	PickableFlag                       *string  `bun:"pickable_flag,type:char(1)"`
	PickLockedFlag                     string   `bun:"pick_locked_flag,type:char(1)"`
	PutableFlag                        *string  `bun:"putable_flag,type:char(1)"`
	PutLockedFlag                      string   `bun:"put_locked_flag,type:char(1)"`
	FrozenFlag                         string   `bun:"frozen_flag,type:char(1)"`
	RfBinFlag                          string   `bun:"rf_bin_flag,type:char(1)"`
	ConsolidationBinFlag               string   `bun:"consolidation_bin_flag,type:char(1)"`
	WeighStationFlag                   *string  `bun:"weigh_station_flag,type:char(1)"`
}
