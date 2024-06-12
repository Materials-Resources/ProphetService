package model

import (
	"github.com/uptrace/bun"
	"time"
)

type InventorySupplierXUom struct {
	bun.BaseModel            `bun:"table:inventory_supplier_x_uom"`
	InventorySupplierXUomUid int32     `bun:"inventory_supplier_x_uom_uid,type:int,pk,identity"`
	ItemUomUid               int32     `bun:"item_uom_uid,type:int"`
	InventorySupplierUid     int32     `bun:"inventory_supplier_uid,type:int"`
	SupplierUnitOfMeasure    string    `bun:"supplier_unit_of_measure,type:varchar(255),nullzero"`
	DateCreated              time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified         time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy         string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	GtinCode                 string    `bun:"gtin_code,type:varchar(255),nullzero"`
}
