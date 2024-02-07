package domain

type InventoryReceipt struct {
	ID    string
	Items []ValidatedInventoryReceiptItem
}

func (e *InventoryReceipt) validate() error { return nil }

type InventoryReceiptItem struct {
	ProductId        string
	ProductName      string
	ProductSn        string
	PrimaryBin       string
	ReceivedQuantity float64
	ReceivedUnit     string
	Orders           []ValidatedInventoryReceiptItemOrder
}

func (e *InventoryReceiptItem) validate() error { return nil }

type InventoryReceiptItemOrder struct {
	ID                string
	CustomerName      string
	AllocatedQuantity float64
	AllocatedUnit     string
}

func (e *InventoryReceiptItemOrder) validate() error { return nil }
