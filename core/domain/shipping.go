package domain

type PickTicket struct {
	ID                 string
	ShippingAddress    PickTicketShippingAddress
	OrderId            string
	OrderPurchaseOrder string
	OrderContactName   string
}

type PickTicketShippingAddress struct {
	Name                 string
	LineOne              string
	LineTwo              string
	City                 string
	State                string
	DeliveryInstructions string
}
