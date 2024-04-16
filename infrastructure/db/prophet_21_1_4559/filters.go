package prophet_21_1_4559

import "github.com/materials-resources/s_prophet/internal/validator"

type Filters struct {
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
	CurrentCursor int
	NextCursor    int
	TotalRecords  int
}

func calculateMetadata(totalRecords, lastRecordCursor, currentRecordCursor int) Metadata {
	if totalRecords == 0 {
		return Metadata{}
	}
	return Metadata{
		TotalRecords:  totalRecords,
		CurrentCursor: currentRecordCursor,
		NextCursor:    lastRecordCursor,
	}
}
