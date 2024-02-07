package domain

import (
	"errors"
)

type Product struct {
	ID          string
	SN          string
	Name        string
	Description string
}

func (p *Product) validate() error {
	if p.Name == "" {
		return errors.New("invalid product details")
	}
	return nil
}

func NewProduct(name string, description string) *Product {
	return &Product{
		Name:        name,
		Description: description,
	}
}
