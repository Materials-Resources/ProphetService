package domain

type Shipment struct {
	Id                   string
	CarrierName          string
	CarrierTracking      *string
	ShippingAddress      Address
	DeliveryInstructions *string
	Contact              Contact
	Customer             Customer
	OrderId              string
	OrderPurchaseOrder   *string
}
