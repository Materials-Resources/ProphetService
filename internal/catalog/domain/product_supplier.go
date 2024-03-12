package domain

import (
	"errors"
)

type ProductSupplier struct {
	ProductId     string
	SupplierId    string
	DivisionId    string
	SupplierSn    string
	ListPrice     float64
	PurchasePrice float64
	Delete        bool
}

type ProductSupplierPatch struct {
	SupplierProductSn *string
	ListPrice         *float64
	PurchasePrice     *float64
	Delete            *bool
}

func (p *ProductSupplier) validate() error {
	if p.ProductId == "" {
		return errors.New("invalid Product Id")
	}
	if p.SupplierId == "" {
		return errors.New("invalid Supplier Id")
	}
	return nil
}

func NewProductSupplier(productId, locationId, supplierId, divisionId, supplierSn string) *ProductSupplier {
	return &ProductSupplier{ProductId: productId, SupplierId: supplierId,
		DivisionId: divisionId, SupplierSn: supplierSn,
	}
}
