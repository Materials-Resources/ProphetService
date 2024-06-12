package model

type Inventoryissuesresults40 struct {
	bun.BaseModel                      `bun:"table:InventoryIssuesResults40"`
	Run                                int32   `bun:"run,type:int"`
	LocationId                         float64 `bun:"location_id,type:decimal(19,0)"`
	ItemId                             string  `bun:"item_id,type:varchar(40)"`
	InvMastUid                         int32   `bun:"inv_mast_uid,type:int"`
	Bin                                string  `bun:"bin,type:varchar(10)"`
	SumTransactionLinesBinQtyAllocated float64 `bun:"sum_transaction_lines_bin_qty_allocated,type:decimal(38,9),nullzero"`
	TableInvBinQtyAllocated            float64 `bun:"table_inv_bin_qty_allocated,type:decimal(19,9)"`
	BinId                              string  `bun:"bin_id,type:varchar(10)"`
	BinType                            string  `bun:"bin_type,type:varchar(255),nullzero"`
	PickableFlag                       string  `bun:"pickable_flag,type:char,nullzero"`
	PickLockedFlag                     string  `bun:"pick_locked_flag,type:char"`
	PutableFlag                        string  `bun:"putable_flag,type:char,nullzero"`
	PutLockedFlag                      string  `bun:"put_locked_flag,type:char"`
	FrozenFlag                         string  `bun:"frozen_flag,type:char"`
	RfBinFlag                          string  `bun:"rf_bin_flag,type:char"`
	ConsolidationBinFlag               string  `bun:"consolidation_bin_flag,type:char"`
	WeighStationFlag                   string  `bun:"weigh_station_flag,type:char,nullzero"`
}
