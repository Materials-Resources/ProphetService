package model

type SupplierTrade struct {
	bun.BaseModel    `bun:"table:supplier_trade"`
	SupplierTradeUid int32     `bun:"supplier_trade_uid,type:int,pk,identity"`
	SupplierId       float64   `bun:"supplier_id,type:decimal(19,0)"`
	NaftaTaxId       string    `bun:"nafta_tax_id,type:varchar(255),nullzero"`
	DateCreated      time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy        string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
