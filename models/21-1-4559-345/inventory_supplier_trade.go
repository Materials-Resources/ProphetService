package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type InventorySupplierTrade struct {
	bun.BaseModel               `bun:"table:inventory_supplier_trade"`
	InventorySupplierTradeUid   int32     `bun:"inventory_supplier_trade_uid,type:int,autoincrement,pk"`       // Unique identifier for this table.
	InventorySupplierUid        int32     `bun:"inventory_supplier_uid,type:int,unique"`                       // Inventory supplier record for which this data pertains.
	CertificateOfOriginFilePath string    `bun:"certificate_of_origin_file_path,type:varchar(255),nullzero"`   // Path for CoO related to the inventory supplier.
	CertificateOfOriginExpDate  time.Time `bun:"certificate_of_origin_exp_date,type:datetime,nullzero"`        // Expiration date tied to this inventory suppliers CoO.
	DateCreated                 time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy                   string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified            time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy            string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	CountryOfOrigin             string    `bun:"country_of_origin,type:varchar(50),nullzero"`                  // The country of origin for the item/supplier.
}
