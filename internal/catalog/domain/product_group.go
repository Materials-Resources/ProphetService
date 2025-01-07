package domain

import (
	"github.com/materials-resources/s-prophet/internal/validator"
)

type ProductGroup struct {
	Sn   string
	Name string
}

type ProductGroupInput struct {
	Sn   string
	Name *string
}

// ValidateUpdate runs validator against fields of ProductGroup
func (i *ProductGroupInput) ValidateUpdate(v *validator.Validator) {
	v.Check(i.Sn != "", "sn", "must be provided")

	if i.Name != nil {
		v.Check(*i.Name != "", "name", "must be provided")
		v.Check(len(*i.Name) <= 80, "name", "must not be more than 80 characters")
	}

}

// ValidateCreate runs validator against fields of ProductGroup
func (i *ProductGroupInput) ValidateCreate(v *validator.Validator) {
	v.Check(i.Sn != "", "sn", "must be provided")
	v.Check(len(i.Sn) <= 8, "sn", "must be more than 8 characters")

	v.Check(i.Name != nil, "name", "must be provided")
	if i.Name != nil {
		v.Check(*i.Name != "", "name", "must be provided")
		v.Check(len(*i.Name) <= 80, "name", "must not be more than 80 characters")
	}
}
