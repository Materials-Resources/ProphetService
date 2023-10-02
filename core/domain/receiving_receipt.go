package domain

type ReceivingReceipt struct {
	Id       string
	Products []*ReceivingReceiptProduct
}

type ReceivingReceiptProduct struct {
	Name            string
	Sn              string
	Id              string
	UnitsReceived   float64
	UnitLabel       string
	PrimaryBin      string
	AllocatedOrders []*ReceivingReceiptAllocatedOrder
}

type ReceivingReceiptAllocatedOrder struct {
	Id             string
	Name           string
	UnitsAllocated float64
	UnitLabel      string
}
