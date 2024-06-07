package domain

import (
	"github.com/materials-resources/s-prophet/internal/validator"
)

type ProductGroup struct {
	Sn   string
	Name *string
}

// ValidateProductGroupUpdate runs validator against fields of ProductGroup
func ValidateProductGroupUpdate(v *validator.Validator, productGroup ProductGroup) {
	v.Check(productGroup.Sn != "", "sn", "must be provided")
	if productGroup.Name != nil {
		v.Check(*productGroup.Name != "", "name", "must be provided")
		v.Check(len(*productGroup.Name) <= 80, "name", "must not be more than 80 characters")
	}
}

// ValidateProductGroupCreate runs validator against fields of ProductGroup
func ValidateProductGroupCreate(v *validator.Validator, productGroup ProductGroup) {
	v.Check(productGroup.Sn != "", "sn", "must be provided")
	v.Check(len(productGroup.Sn) <= 8, "sn", "must be more than 8 characters")
	v.Check(productGroup.Name != nil && *productGroup.Name != "", "name", "must be provided")
	v.Check(productGroup.Name == nil || len(*productGroup.Name) <= 80, "name", "must not be more than 80 characters")
}
