package data

import (
	"context"
	"fmt"
	prophet "github.com/materials-resources/s-prophet/models/21-1-4559-345"
	"github.com/uptrace/bun"
)

type AlternateCode struct {
	prophet.AlternateCode `bun:",extend"`
}

type AverageInventoryValue struct {
	prophet.AverageInventoryValue `bun:",extend"`
}

type AssemblyHdr struct {
	prophet.AssemblyHdr `bun:",extend"`

	AssemblyLines []*AssemblyLine `bun:"rel:has-many,join:inv_mast_uid=inv_mast_uid"`
}

type AssemblyLine struct {
	prophet.AssemblyLine `bun:",extend"`
}

type PricePage struct {
	prophet.PricePage `bun:",extend"`

	PricePageXBooks []*PricePageXBook `bun:"rel:has-many,join:price_page_uid=price_page_uid"`
}

type PricePageXBook struct {
	prophet.PricePageXBook `bun:",extend"`
}

type InvAdjLine struct {
	prophet.InvAdjLine `bun:",extend"`
}

type InvBin struct {
	prophet.InvBin `bun:",extend"`
}

type InvLoc struct {
	prophet.InvLoc      `bun:",extend"`
	InvMast             *InvMast             `bun:"rel:belongs-to,join:inv_mast_uid=inv_mast_uid"`
	ProductGroup        *ProductGroup        `bun:"rel:has-one,join:product_group_id=product_group_id"`
	InvBins             []*InvBin            `bun:"rel:has-many,join:company_id=company_id,join:location_id=location_id,join:inv_mast_uid=inv_mast_uid"`
	InvTrans            []*InvTran           `bun:"rel:has-many,join:location_id=location_id,join:inv_mast_uid=inv_mast_uid"`
	InvLocStockStatuses []*InvLocStockStatus `bun:"rel:has-many,join:location_id=location_id,join:inv_mast_uid=inv_mast_uid"`
}

type InvLocMsp struct {
	prophet.InvLocMsp `bun:",extend"`
}

type InvLocStockStatus struct {
	prophet.InvLocStockStatus `bun:",extend"`
}

type InvMast struct {
	prophet.InvMast `bun:",extend"`

	AssemblyHdr *AssemblyHdr `bun:"rel:has-one,join:inv_mast_uid=inv_mast_uid"`

	AverageInventoryValue []*AverageInventoryValue `bun:"rel:has-many,join:inv_mast_uid=inv_mast_uid"`
	AlternateCodes        []*AlternateCode         `bun:"rel:has-many,join:inv_mast_uid=inv_mast_uid"`
	ItemIdChangeHistory   []*ItemIdChangeHistory   `bun:"rel:has-many,join:inv_mast_uid=inv_mast_uid"`
	InvAdjLines           []*InvAdjLine            `bun:"rel:has-many,join:inv_mast_uid=inv_mast_uid"`
	InvLocs               []*InvLoc                `bun:"rel:has-many,join:inv_mast_uid=inv_mast_uid"`
	InvLocMsps            []*InvLocMsp             `bun:"rel:has-many,join:inv_mast_uid=inv_mast_uid"`
	InventorySuppliers    []*InventorySupplier     `bun:"rel:has-many,join:inv_mast_uid=inv_mast_uid"`
	ItemCategoryXInvMasts []*ItemCategoryXInvMast  `bun:"rel:has-many,join:inv_mast_uid=inv_mast_uid"`
	ItemConversions       []*ItemConversion        `bun:"rel:has-many,join:inv_mast_uid=inv_mast_uid"`
	ItemUoms              []*ItemUom               `bun:"rel:has-many,join:inv_mast_uid=inv_mast_uid"`
	PricePages            []*PricePage             `bun:"rel:has-many,join:inv_mast_uid=inv_mast_uid"`
}

type InvTran struct {
	prophet.InvTran `bun:",extend"`
}

type InventoryReceiptsLine struct {
	prophet.InventoryReceiptsLine `bun:",extend"`
}

type InventorySupplier struct {
	prophet.InventorySupplier `bun:",extend"`

	InventorySupplierXLocs []*InventorySupplierXLoc `bun:"rel:has-many,join:inventory_supplier_uid=inventory_supplier_uid"`
	InventorySupplierXUoms []*InventorySupplierXUom `bun:"rel:has-many,join:inventory_supplier_uid=inventory_supplier_uid"`
}

type InventorySupplierXLoc struct {
	prophet.InventorySupplierXLoc `bun:",extend"`
}

type InventorySupplierXUom struct {
	prophet.InventorySupplierXUom `bun:",extend"`
}

type ItemCategoryXInvMast struct {
	prophet.ItemCategoryXInvMast `bun:",extend"`
}

type ItemConversion struct {
	prophet.ItemConversion `bun:",extend"`
}

type ItemIdChangeHistory struct {
	prophet.ItemIdChangeHistory `bun:",extend"`
}

type ItemUom struct {
	prophet.ItemUom `bun:",extend"`
}

type OeLine struct {
	prophet.OeLine `bun:",extend"`
}

type ProductGroup struct {
	prophet.ProductGroup `bun:",extend"`
}

type Models struct {
	bun bun.IDB

	InvLoc       InvLocModel
	InvMast      InvMastModel
	ProductGroup ProductGroupModel
}

func NewModels(bun bun.IDB) *Models {
	return &Models{
		bun:          bun,
		InvLoc:       NewInvLocModel(bun),
		InvMast:      NewInvMastModel(bun),
		ProductGroup: NewProductGroupModel(bun),
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
