package domain

type Product struct {
	Name           string
	Description    string
	SN             string
	ID             int32
	Price          float64
	UnitsAvailable float64
	UnitLabel      string
}

type ProductFilter struct {
	Cursor Cursor
	Id     []int32
}
