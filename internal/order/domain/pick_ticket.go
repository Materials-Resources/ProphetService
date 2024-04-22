package domain

type PickTicket struct {
	Id                   float64
	ShippingAddress      Address
	DeliveryInstructions string
	OrderId              string
	OrderPurchaseOrder   string
	OrderContact         Contact
}
