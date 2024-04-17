package prophet_21_1_4559

import (
	"context"
	"database/sql"
	"github.com/uptrace/bun"
	"time"
)

type InventorySupplierXUom struct {
	bun.BaseModel            `bun:"table:inventory_supplier_x_uom"`
	InventorySupplierXUomUid int32          `bun:"inventory_supplier_x_uom_uid,type:int,pk,identity"`
	ItemUomUid               int32          `bun:"item_uom_uid,type:int"`
	InventorySupplierUid     int32          `bun:"inventory_supplier_uid,type:int"`
	SupplierUnitOfMeasure    sql.NullString `bun:"supplier_unit_of_measure,type:varchar(255),nullzero"`
	DateCreated              time.Time      `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                string         `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified         time.Time      `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy         string         `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
	GtinCode                 sql.NullString `bun:"gtin_code,type:varchar(255),nullzero"`
}

type InventorySupplierXUomModel struct {
	bun bun.IDB
}

// GetByItemUomUid returns an slice of InventorySupplierXUom by the given ItemUomUid
func (m InventorySupplierXUomModel) GetByItemUomUid(ctx context.Context, itemUomUid int32) (
	[]*InventorySupplierXUom,
	error) {
	var inventorySupplierXUoms []*InventorySupplierXUom
	err := m.bun.NewSelect().Model(&inventorySupplierXUoms).Where("item_uom_uid = ?", itemUomUid).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return inventorySupplierXUoms, nil
}

// Delete deletes the InventorySupplierXUom from the database.
func (m InventorySupplierXUomModel) Delete(
	ctx context.Context, inventorySupplierXUom *InventorySupplierXUom) error {
	_, err := m.bun.NewDelete().Model(inventorySupplierXUom).WherePK().Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
