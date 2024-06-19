package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type InventorySupplierXUom struct {
	bun.BaseModel            `bun:"table:inventory_supplier_x_uom"`
	InventorySupplierXUomUid int32     `bun:"inventory_supplier_x_uom_uid,type:int,autoincrement,pk"`       // Table uinque identifier
	ItemUomUid               int32     `bun:"item_uom_uid,type:int,unique"`                                 // Unique identifier of the Unit of Measure record
	InventorySupplierUid     int32     `bun:"inventory_supplier_uid,type:int,unique"`                       // Unique identifier of the Supplier record
	SupplierUnitOfMeasure    string    `bun:"supplier_unit_of_measure,type:varchar(255),nullzero"`          // Unit of Measure Alias used by the Supplier
	DateCreated              time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Date and time the record was originally created
	CreatedBy                string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`         // User who created the record
	DateLastModified         time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Date and time the record was modified
	LastMaintainedBy         string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"` // User who last changed the record
	GtinCode                 string    `bun:"gtin_code,type:varchar(255),nullzero"`                         // Use to store the gtin code
}
