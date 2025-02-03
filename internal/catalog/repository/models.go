package repository

import (
	"context"
	"fmt"
	prophet "github.com/materials-resources/s-prophet/models/21-1-4559-345"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/schema"
	"time"
)

type alternateCode struct {
	prophet.AlternateCode `bun:",extend"`
}

type averageInventoryValue struct {
	prophet.AverageInventoryValue `bun:",extend"`
}

type assemblyHdr struct {
	prophet.AssemblyHdr `bun:",extend"`

	AssemblyLines []*assemblyLine `bun:"rel:has-many,join:inv_mast_uid=inv_mast_uid"`
}

type assemblyLine struct {
	prophet.AssemblyLine `bun:",extend"`
}

type pricePage struct {
	prophet.PricePage `bun:",extend"`

	PricePageXBooks []*pricePageXBook `bun:"rel:has-many,join:price_page_uid=price_page_uid"`
}

type pricePageXBook struct {
	prophet.PricePageXBook `bun:",extend"`
}

type invAdjLine struct {
	prophet.InvAdjLine `bun:",extend"`
}

type invBin struct {
	prophet.InvBin `bun:",extend"`
}

type invLoc struct {
	prophet.InvLoc      `bun:",extend"`
	InvMast             *InvMast             `bun:"rel:belongs-to,join:inv_mast_uid=inv_mast_uid"`
	ProductGroup        *productGroup        `bun:"rel:has-one,join:product_group_id=product_group_id"`
	InvBins             []*invBin            `bun:"rel:has-many,join:company_id=company_id,join:location_id=location_id,join:inv_mast_uid=inv_mast_uid"`
	InvTrans            []*invTran           `bun:"rel:has-many,join:inv_mast_uid=inv_mast_uid,join:location_id=location_id"`
	InvLocStockStatuses []*invLocStockStatus `bun:"rel:has-many,join:inv_mast_uid=inv_mast_uid,join:location_id=location_id"`
}

type invLocMsp struct {
	prophet.InvLocMsp `bun:",extend"`
}

type invLocStockStatus struct {
	prophet.InvLocStockStatus `bun:",extend"`
}

type invSub struct {
	prophet.InvSub `bun:",extend"`
}

type invTran struct {
	prophet.InvTran `bun:",extend"`
}

type invXref struct {
	prophet.InvXref `bun:",extend"`
}

type inventoryReceiptsLine struct {
	prophet.InventoryReceiptsLine `bun:",extend"`
}

type inventorySupplier struct {
	prophet.InventorySupplier `bun:",extend"`

	InventorySupplierXLocs []*inventorySupplierXLoc `bun:"rel:has-many,join:inventory_supplier_uid=inventory_supplier_uid"`
	InventorySupplierXUoms []*inventorySupplierXUom `bun:"rel:has-many,join:inventory_supplier_uid=inventory_supplier_uid"`
}

type inventorySupplierXLoc struct {
	prophet.InventorySupplierXLoc `bun:",extend"`
}

type inventorySupplierXUom struct {
	prophet.InventorySupplierXUom `bun:",extend"`
}

type itemCategoryXInvMast struct {
	prophet.ItemCategoryXInvMast `bun:",extend"`
}

type itemConversion struct {
	prophet.ItemConversion `bun:",extend"`
}

type itemIdChangeHistory struct {
	prophet.ItemIdChangeHistory `bun:",extend"`
}

type itemUom struct {
	prophet.ItemUom `bun:",extend"`
}
type productGroup struct {
	prophet.ProductGroup `bun:",extend"`
}

func (m *productGroup) BeforeAppendModel(ctx context.Context, query schema.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		m.DateCreated = time.Now()
		m.DateLastModified = time.Now()
	case *bun.UpdateQuery:
		m.DateLastModified = time.Now()
	}
	return nil
}

type Models struct {
	bun bun.IDB

	InvLoc InvLocModel
}

func NewModels(bun bun.IDB) *Models {
	return &Models{
		bun:    bun,
		InvLoc: NewInvLocModel(bun),
	}
}

func (m *Models) ModelsTx(ctx context.Context, fn func(models *Models) error) error {
	tx, err := m.bun.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := NewModels(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}
