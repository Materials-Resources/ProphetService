package model

type TpcxDeadStock struct {
	bun.BaseModel         `bun:"table:tpcx_dead_stock"`
	TpcxDeadStockUid      int32     `bun:"tpcx_dead_stock_uid,type:int,pk"`
	TpcxTradingPartnerUid int32     `bun:"tpcx_trading_partner_uid,type:int"`
	InvMastUid            int32     `bun:"inv_mast_uid,type:int"`
	TpcxItemDesc          string    `bun:"tpcx_item_desc,type:varchar(255),nullzero"`
	TpcxUpcCd             string    `bun:"tpcx_upc_cd,type:varchar(255),nullzero"`
	TpcxQtyAvailable      float64   `bun:"tpcx_qty_available,type:decimal(19,9)"`
	TpcxPrice             float64   `bun:"tpcx_price,type:decimal(19,9)"`
	DateCreated           time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy             string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified      time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy      string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
