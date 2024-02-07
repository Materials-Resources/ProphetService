package domain

type ValidatedInventoryReceipt struct {
	InventoryReceipt
	isValidated bool
}

func (v *ValidatedInventoryReceipt) IsValid() bool { return v.isValidated }

func NewValidatedInventoryReceipt(inventoryReceipt *InventoryReceipt) (*ValidatedInventoryReceipt, error) {
	if err := inventoryReceipt.validate(); err != nil {
		return nil, err
	}
	return &ValidatedInventoryReceipt{InventoryReceipt: *inventoryReceipt, isValidated: true}, nil
}

type ValidatedInventoryReceiptItem struct {
	InventoryReceiptItem
	isValidated bool
}

func (v *ValidatedInventoryReceiptItem) IsValid() bool { return v.isValidated }

func NewValidatedInventoryReceiptItem(item *InventoryReceiptItem) (*ValidatedInventoryReceiptItem, error) {
	if err := item.validate(); err != nil {
		return nil, err
	}
	return &ValidatedInventoryReceiptItem{InventoryReceiptItem: *item, isValidated: true}, nil
}

type ValidatedInventoryReceiptItemOrder struct {
	InventoryReceiptItemOrder
	isValidated bool
}

func (v *ValidatedInventoryReceiptItemOrder) IsValid() bool {
	return v.isValidated
}

func NewValidatedInventoryReceiptItemOrder(order *InventoryReceiptItemOrder) (*ValidatedInventoryReceiptItemOrder,
	error,
) {
	if err := order.validate(); err != nil {
		return nil, err
	}
	return &ValidatedInventoryReceiptItemOrder{InventoryReceiptItemOrder: *order, isValidated: true}, nil
}
