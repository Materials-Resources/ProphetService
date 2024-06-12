package model

type InventorySupplierTrade struct {
	bun.BaseModel               `bun:"table:inventory_supplier_trade"`
	InventorySupplierTradeUid   int32     `bun:"inventory_supplier_trade_uid,type:int,pk,identity"`
	InventorySupplierUid        int32     `bun:"inventory_supplier_uid,type:int"`
	CertificateOfOriginFilePath string    `bun:"certificate_of_origin_file_path,type:varchar(255),nullzero"`
	CertificateOfOriginExpDate  time.Time `bun:"certificate_of_origin_exp_date,type:datetime,nullzero"`
	DateCreated                 time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                   string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified            time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy            string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	CountryOfOrigin             string    `bun:"country_of_origin,type:varchar(50),nullzero"`
}
