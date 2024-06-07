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
	Id                        string
	CarrierName               string
	CarrierTracking           string
	ShippingAddressId         string
	ShippingAddressName       string
	ShippingAddressLineOne    string
	ShippingAddressLineTwo    string
	ShippingAddressCity       string
	ShippingAddressState      string
	ShippingAddressPostalCode string
	ShippingAddressCountry    string
	DeliveryInstructions      string
	ContactId                 string
	ContactName               string
	OrderId                   string
	OrderPurchaseOrder        string
}
