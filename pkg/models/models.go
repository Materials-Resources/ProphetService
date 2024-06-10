package models

import (
	"errors"
	"github.com/uptrace/bun"
)

var (
	ErrRecordNotFound  = errors.New("could not find requested record")
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
	Address               AddressModel
	AlternateCode         AlternateCodeModel
	AssemblyHdr           AssemblyHdrModel
	AssemblyLine          AssemblyLineModel
	AverageInventoryValue AverageInventoryValueModel
	Contacts              ContactsModel
	ContactsXShipTo       ContactsXShipToModel
	Customer              CustomerModel
	InvBin                InvBinModel
	InvMast               InvMastModel
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
	InventoryReceiptsHdr  InventoryReceiptsHdrModel
	InventoryReceiptsLine InventoryReceiptsLineModel
	InvoiceHdr            InvoiceHdrModel
	ItemConversion        ItemConversionModel
	ItemCategoryXInvMast  ItemCategoryXInvMastModel
	OeHdr                 OeHdrModel
	OeHdrSalesrep         OeHdrSalesrepModel
	OeLine                OeLineModel
	OePickTicket          OePickTicketModel
	PricePage             PricePageModel
	PricePageXBook        PricePageXBookModel
	ProductGroup          ProductGroupModel
	QuoteHdr              QuoteHdrModel
	QuoteLine             QuoteLineModel
	ShipTo                ShipToModel
	ShipToSalesrep        ShipToSalesrepModel
}

func NewModels(db bun.IDB) *Models {
	return &Models{
		Bun:                   db,
		Address:               AddressModel{bun: db},
		AlternateCode:         AlternateCodeModel{bun: db},
		AssemblyHdr:           AssemblyHdrModel{bun: db},
		AssemblyLine:          AssemblyLineModel{bun: db},
		AverageInventoryValue: AverageInventoryValueModel{bun: db},
		Contacts:              ContactsModel{bun: db},
		ContactsXShipTo:       ContactsXShipToModel{bun: db},
		Customer:              CustomerModel{bun: db},
		InvBin:                InvBinModel{bun: db},
		InvMast:               InvMastModel{bun: db},
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
		InventoryReceiptsHdr:  InventoryReceiptsHdrModel{bun: db},
		InventoryReceiptsLine: InventoryReceiptsLineModel{bun: db},
		InvoiceHdr:            InvoiceHdrModel{bun: db},
		ItemConversion:        ItemConversionModel{bun: db},
		ItemCategoryXInvMast:  ItemCategoryXInvMastModel{bun: db},
		OeHdr:                 OeHdrModel{bun: db},
		OeHdrSalesrep:         OeHdrSalesrepModel{bun: db},
		OeLine:                OeLineModel{bun: db},
		OePickTicket:          OePickTicketModel{bun: db},
		PricePage:             PricePageModel{bun: db},
		PricePageXBook:        PricePageXBookModel{bun: db},
		ProductGroup:          ProductGroupModel{bun: db},
		QuoteHdr:              QuoteHdrModel{bun: db},
		QuoteLine:             QuoteLineModel{bun: db},
		ShipTo:                ShipToModel{bun: db},
		ShipToSalesrep:        ShipToSalesrepModel{bun: db},
	}

}
