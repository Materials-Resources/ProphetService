package domain

type Purchase struct {
	ProductId          string
	ProductSn          *string
	ProductName        *string
	ProductDescription *string
	UnitType           *string
	QuantityPurchased  *float64
}
