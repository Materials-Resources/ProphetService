package domain

type ProductFilter struct {
	ProductGroupSn string
	ProductID      string
	Limit          int
	Cursor         int32
}
