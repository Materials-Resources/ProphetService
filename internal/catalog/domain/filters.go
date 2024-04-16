package domain

import "github.com/materials-resources/s_prophet/internal/validator"

type Filter struct {
	Limit  int
	Cursor int
}

type Metadata struct {
	CurrentCursor int
	NextCursor    int
	Limit         int
}

// ValidateFilter validates the filter.
func ValidateFilter(v *validator.Validator, f Filter) {
	v.Check(f.Limit > 0, "limit", "must be greater than zero")
	v.Check(f.Limit <= 500, "limit", "must be a maximum of 500")

	v.Check(f.Cursor >= 0, "cursor", "must be greater than or equal to zero")
}
