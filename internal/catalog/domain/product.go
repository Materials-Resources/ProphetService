package domain

import (
	"github.com/materials-resources/s-prophet/internal/validator"
)

type Product struct {
	Uid              string
	Sn               string
	Name             string
	Description      string
	ProductGroupId   string
	ProductGroupName string
	ListPrice        float64
	StockQuantity    float64
}

type ProductPatch struct {
}

func ValidateProduct(v *validator.Validator, product Product) {
	v.Check(product.Name != "", "name", "must be provided")
	v.Check(len(product.Name) <= 40, "name", "must not be more than 40 characters")

	v.Check(product.Description != "", "description", "must be provided")
	v.Check(len(product.Description) <= 255, "description", "must not be more than 255 characters")
}
