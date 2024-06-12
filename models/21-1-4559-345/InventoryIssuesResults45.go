package model

type Inventoryissuesresults45 struct {
	bun.BaseModel       `bun:"table:InventoryIssuesResults45"`
	Run                 int32   `bun:"run,type:int"`
	TranType            string  `bun:"tran_type,type:varchar(255),nullzero"`
	ItemId              string  `bun:"item_id,type:varchar(40),nullzero"`
	LocationId          float64 `bun:"location_id,type:decimal(19,0),nullzero"`
	InvMastUid          int32   `bun:"inv_mast_uid,type:int,nullzero"`
	DocumentNo          float64 `bun:"document_no,type:decimal(19,0),nullzero"`
	LineNo              int32   `bun:"line_no,type:int,nullzero"`
	BinCd               string  `bun:"bin_cd,type:varchar(40),nullzero"`
	TranBinQtyAllocated float64 `bun:"tran_bin_qty_allocated,type:decimal(20,13),nullzero"`
	SubLineNo           int32   `bun:"sub_line_no,type:int,nullzero"`
	DocumentLineBinUid  int32   `bun:"document_line_bin_uid,type:int,nullzero"`
}
