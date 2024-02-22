package domain

import (
	"errors"
)

type ProductGroup struct {
	Name string
	SN   string
}

func (p *ProductGroup) validate() error {
	if p.Name == "" {
		return errors.New("invalid product group details")
	}
	return nil
}

func NewProductGroup(name, sn string) *ProductGroup {
	return &ProductGroup{
		Name: name,
		SN:   sn,
	}
}
