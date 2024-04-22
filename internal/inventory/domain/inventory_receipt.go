package domain

type InventoryReceipt struct {
	ID    float64
	Items []InventoryReceiptItem
}

type InventoryReceiptItem struct {
	ProductUid       int32
	PrimaryBin       string
	ReceivedQuantity float64
	ReceivedUnit     string
	AllocatedOrders  []string
}
