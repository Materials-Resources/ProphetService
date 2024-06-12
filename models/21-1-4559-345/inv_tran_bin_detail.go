package model

type InvTranBinDetail struct {
	bun.BaseModel       `bun:"table:inv_tran_bin_detail"`
	CompanyId           string    `bun:"company_id,type:varchar(8)"`
	LocationId          float64   `bun:"location_id,type:decimal(19,0)"`
	Bin                 string    `bun:"bin,type:varchar(10)"`
	TransactionNumber   int32     `bun:"transaction_number,type:int,default:((-1))"`
	Quantity            float64   `bun:"quantity,type:decimal(19,9)"`
	DateCreated         time.Time `bun:"date_created,type:datetime"`
	DateLastModified    time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy    string    `bun:"last_maintained_by,type:varchar(30)"`
	InvTranBinDetailUid int32     `bun:"inv_tran_bin_detail_uid,type:int,pk"`
	DocumentLineBinUid  int32     `bun:"document_line_bin_uid,type:int"`
	InvMastUid          int32     `bun:"inv_mast_uid,type:int"`
	QtyAllocated        float64   `bun:"qty_allocated,type:decimal(19,9),default:(0)"`
}
