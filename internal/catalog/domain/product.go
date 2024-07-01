package domain

import (
	"github.com/materials-resources/s-prophet/internal/validator"
)

type Product struct {
	Uid              string
	Sn               *string
	Name             *string
	Description      *string
	ProductGroupSn   *string
	ProductGroupName string
	StockQuantity    float64
}

type ProductFilters struct {
	ProductGroupSn *[]string
}

// ValidateProductCreate validates the product creation request.
func ValidateProductCreate(v *validator.Validator, product Product) {
	v.Check(product.Sn != nil && *product.Sn != "", "sn", "must be provided")
	v.Check(product.Sn == nil || len(*product.Sn) <= 40, "sn", "must not be more than 40 characters")

	v.Check(product.Name != nil && *product.Name != "", "name", "must be provided")
	v.Check(product.Name == nil || len(*product.Name) <= 40, "name", "must not be more than 40 characters")

	v.Check(product.Description != nil && *product.Description != "", "description", "must be provided")
	v.Check(product.Description == nil || len(*product.Description) <= 255, "description", "must not be more than 255 characters")

	v.Check(product.ProductGroupSn != nil && *product.ProductGroupSn != "", "product_group_sn", "must be provided")
}

// ValidateProductUpdate validates the product update request.
func ValidateProductUpdate(v *validator.Validator, product Product) {
	v.Check(product.Uid != "", "uid", "must be provided")

	if product.Sn != nil {
		v.Check(*product.Sn != "", "sn", "must be provided")
		v.Check(len(*product.Sn) <= 40, "sn", "must not be more than 40 characters")
	}

	if product.Name != nil {
		v.Check(*product.Name != "", "name", "must be provided")
		v.Check(len(*product.Name) <= 40, "name", "must not be more than 40 characters")
	}

	if product.Description != nil {
		v.Check(*product.Description != "", "description", "must be provided")
		v.Check(len(*product.Description) <= 255, "description", "must not be more than 255 characters")
	}

	if product.ProductGroupSn != nil {
		v.Check(*product.ProductGroupSn != "", "product_group_sn", "must be provided")

	}
}
