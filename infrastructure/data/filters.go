package data

import "github.com/materials-resources/s_prophet/internal/validator"

type Direction int32

const (
	Next Direction = iota
	Previous
)

type Filters struct {
	Direction    Direction
	Limit        int32
	Cursor       int32
	Sort         string
	SortSafeList []string
}

func ValidateFilters(v *validator.Validator, f Filters) {
	v.Check(f.Limit > 0, "page_size", "must be greater than zero")
	v.Check(f.Cursor >= 0, "cursor", "must be greater than or equal to zero")
}

func (f Filters) cursor() int32 {
	return f.Cursor
}

func (f Filters) limit() int32 {
	return f.Limit
}

type Metadata struct {
	NextCursor     int32
	PreviousCursor int32
	TotalRecords   int32
}

func calculateMetadata(totalRecords, nextCursor, previousCursor int32) Metadata {
	if totalRecords == 0 {
		return Metadata{}
	}
	return Metadata{
		TotalRecords:   totalRecords,
		PreviousCursor: nextCursor,
		NextCursor:     previousCursor,
	}
}
