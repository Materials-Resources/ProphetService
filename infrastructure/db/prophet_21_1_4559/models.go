package prophet_21_1_4559

import (
	"errors"
	"github.com/uptrace/bun"
)

var (
	ErrNotFound        = errors.New("could not find requested record")
	ErrDuplicateRecord = errors.New("a record with this key already exists")
)

type DeleteFlag int

const (
	Undefined DeleteFlag = iota
	DeleteFlagYes
	DeleteFlagNo
)

func (f DeleteFlag) String() string {
	return [...]string{"", "Y", "N"}[f]
}

type Models struct {
	Bun                   bun.IDB
	AlternateCode         AlternateCodeModel
	AverageInventoryValue AverageInventoryValueModel
	InvBin                InvBinModel
	InvLocStockStatus     InvLocStockStatusModel
	InvMast               InvMastModel
	InvLoc                InvLocModel
	ItemUom               ItemUomModel
	InventorySupplier     InventorySupplierModel
	InventorySupplierXLoc InventorySupplierXLocModel
	InventorySupplierXUom InventorySupplierXUomModel
	OeHdr                 OeHdrModel
	ProductGroup          ProductGroupModel
}

func NewModels(db bun.IDB) *Models {
	return &Models{
		Bun:                   db,
		AlternateCode:         AlternateCodeModel{bun: db},
		AverageInventoryValue: AverageInventoryValueModel{bun: db},
		InvBin:                InvBinModel{bun: db},
		InvLocStockStatus:     InvLocStockStatusModel{bun: db},
		InvMast:               InvMastModel{bun: db},
		InvLoc:                InvLocModel{bun: db},
		ItemUom:               ItemUomModel{bun: db},
		InventorySupplier:     InventorySupplierModel{bun: db},
		InventorySupplierXLoc: InventorySupplierXLocModel{bun: db},
		InventorySupplierXUom: InventorySupplierXUomModel{bun: db},
		OeHdr:                 OeHdrModel{bun: db},
		ProductGroup:          ProductGroupModel{bun: db},
	}

}
