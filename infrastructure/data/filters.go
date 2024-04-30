package data

import "github.com/materials-resources/s_prophet/internal/validator"

type PageDirection int32

const (
	PageDirectionUnknown PageDirection = iota
	PageDirectionNext
	PageDirectionPrevious
)

type Filters struct {
	Direction    PageDirection
	Limit        int
	Cursor       int
	Sort         string
	SortSafeList []string
}

func ValidateFilters(v *validator.Validator, f Filters) {
	v.Check(f.Limit > 0, "page_size", "must be greater than zero")
	v.Check(f.Cursor >= 0, "cursor", "must be greater than or equal to zero")
}

func (f Filters) cursor() int {
	return f.Cursor
}

func (f Filters) limit() int {
	return f.Limit
}

type Metadata struct {
	NextCursor     int
	PreviousCursor int
	TotalRecords   int
}

func calculateMetadata(totalRecords, nextCursor, previousCursor int) Metadata {
	if totalRecords == 0 {
		return Metadata{}
	}
	return Metadata{
		TotalRecords:   totalRecords,
		PreviousCursor: nextCursor,
		NextCursor:     previousCursor,
	}
}
