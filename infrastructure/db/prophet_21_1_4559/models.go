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
	AssemblyHdr           AssemblyHdrModel
	AssemblyLine          AssemblyLineModel
	AverageInventoryValue AverageInventoryValueModel
	InvBin                InvBinModel
	InvMast               InvMastModel
	InvLoc                InvLocModel
	InvLocMsp             InvLocMspModel
	InvLocStockStatus     InvLocStockStatusModel
	ItemIdChangeHistory   ItemIdChangeHistoryModel
	ItemUom               ItemUomModel
	InvAdjHdr             InvAdjHdrModel
	InvAdjLine            InvAdjLineModel
	InvTran               InvTranModel
	InventorySupplier     InventorySupplierModel
	InventorySupplierXLoc InventorySupplierXLocModel
	InventorySupplierXUom InventorySupplierXUomModel
	InventoryReceiptsLine InventoryReceiptsLineModel
	ItemConversion        ItemConversionModel
	ItemCategoryXInvMast  ItemCategoryXInvMastModel
	OeHdr                 OeHdrModel
	OeLine                OeLineModel
	PricePage             PricePageModel
	PricePageXBook        PricePageXBookModel
	ProductGroup          ProductGroupModel
}

func NewModels(db bun.IDB) *Models {
	return &Models{
		Bun:                   db,
		AlternateCode:         AlternateCodeModel{bun: db},
		AssemblyHdr:           AssemblyHdrModel{bun: db},
		AssemblyLine:          AssemblyLineModel{bun: db},
		AverageInventoryValue: AverageInventoryValueModel{bun: db},
		InvBin:                InvBinModel{bun: db},
		InvMast:               InvMastModel{bun: db},
		InvLoc:                InvLocModel{bun: db},
		InvLocMsp:             InvLocMspModel{bun: db},
		ItemIdChangeHistory:   ItemIdChangeHistoryModel{bun: db},
		InvLocStockStatus:     InvLocStockStatusModel{bun: db},
		ItemUom:               ItemUomModel{bun: db},
		InvAdjHdr:             InvAdjHdrModel{bun: db},
		InvAdjLine:            InvAdjLineModel{bun: db},
		InvTran:               InvTranModel{bun: db},
		InventorySupplier:     InventorySupplierModel{bun: db},
		InventorySupplierXLoc: InventorySupplierXLocModel{bun: db},
		InventorySupplierXUom: InventorySupplierXUomModel{bun: db},
		InventoryReceiptsLine: InventoryReceiptsLineModel{bun: db},
		ItemConversion:        ItemConversionModel{bun: db},
		ItemCategoryXInvMast:  ItemCategoryXInvMastModel{bun: db},
		OeHdr:                 OeHdrModel{bun: db},
		OeLine:                OeLineModel{bun: db},
		PricePage:             PricePageModel{bun: db},
		PricePageXBook:        PricePageXBookModel{bun: db},
		ProductGroup:          ProductGroupModel{bun: db},
	}

}
