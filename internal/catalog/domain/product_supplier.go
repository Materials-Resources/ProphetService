package domain

import (
	"github.com/materials-resources/s_prophet/internal/validator"
)

type ProductSupplier struct {
	ProductUid        int32
	SupplierId        float64
	DivisionId        float64
	SupplierProductSn string
	ListPrice         float64
	PurchasePrice     float64
	Delete            bool
}

type ProductSupplierPatch struct {
	SupplierProductSn *string
	ListPrice         *float64
	PurchasePrice     *float64
	Delete            *bool
}

func ValidateProductSupplier(v *validator.Validator, productSupplier ProductSupplier) {
	v.Check(productSupplier.ProductUid != 0, "product_uid", "must be provided")
	v.Check(productSupplier.SupplierId != 0, "supplier_id", "must be provided")
}
