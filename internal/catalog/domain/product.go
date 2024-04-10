package domain

import (
	"errors"
)

type Product struct {
	ID               int32
	SN               string
	Name             string
	Description      string
	ProductGroup     *ProductGroup
	ProductGroupId   string
	ProductGroupName string
	ListPrice        float64
	StockQuantity    float64
}

type ProductPatch struct {
}

func (p *Product) validate() error {
	if p.Name == "" {
		return errors.New("invalid product details")
	}
	return nil
}

func NewProduct(name, description string) *Product {
	return &Product{
		Name:        name,
		Description: description,
	}
}
