package domain

type PickTicket struct {
	Id                   float64
	ShippingAddress      Address
	DeliveryInstructions string
	OrderId              string
	OrderPurchaseOrder   string
	OrderContact         Contact
}

type Shipment struct {
	Id              string
	OrderId         string
	InvoiceId       string
	CarrierId       string
	CarrierName     string
	CarrierTracking string
}
