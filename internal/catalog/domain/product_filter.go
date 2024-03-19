package domain

type ProductFilter struct {
	GroupID   string
	ProductID string
	Limit     int
	Cursor    int32
}
