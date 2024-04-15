package domain

import (
	"github.com/materials-resources/s_prophet/internal/validator"
)

type ProductGroup struct {
	UID  int32
	SN   string
	Name string
}

// ValidateProductGroup runs validator against fields of ProductGroup
func ValidateProductGroup(v *validator.Validator, productGroup ProductGroup) {
	v.Check(productGroup.Name != "", "name", "must be provided")
	v.Check(len(productGroup.Name) <= 80, "name", "must not be more than 80 characters")

	v.Check(productGroup.SN != "", "sn", "must be provided")
	v.Check(len(productGroup.SN) <= 8, "sn", "must be more than 8 characters")
}
