package domain

type Order struct {
	ID              string
	PurchaseOrder   string
	OrderItems      []OrderItem
	ShippingAddress OrderShippingAddress
}
type OrderShippingAddress struct {
	Name                 string
	LineOne              string
	LineTwo              string
	City                 string
	State                string
	DeliveryInstructions string
}

type OrderItem struct {
	ID            string
	Name          string
	UnitPurchased int64
	UnitLabel     string
	CostPerUnit   int64
	SN            string
}
