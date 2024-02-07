package domain

type ValidatedOrder struct {
	Order
	isValidated bool
}

func (v *ValidatedOrder) IsValid() bool {
	return v.isValidated
}

func NewValidatedOrder(order *Order) (*ValidatedOrder, error) {
	if err := order.validate(); err != nil {
		return nil, err
	}
	return &ValidatedOrder{Order: *order, isValidated: true}, nil
}

type ValidatedOrderItem struct {
	OrderItem
	isValidated bool
}

func (v *ValidatedOrderItem) IsValid() bool { return v.isValidated }

func NewValidatedOrderItem(orderItem *OrderItem) (*ValidatedOrderItem, error) {
	if err := orderItem.validate(); err != nil {
		return nil, err
	}
	return &ValidatedOrderItem{OrderItem: *orderItem, isValidated: true}, nil
}

type ValidatedOrderCustomerContact struct {
	OrderCustomerContact
	isValidated bool
}

func (v *ValidatedOrderCustomerContact) IsValid() bool { return v.isValidated }

func NewValidatedOrderCustomerContact(orderCustomerContact *OrderCustomerContact) (*ValidatedOrderCustomerContact,
	error,
) {
	if err := orderCustomerContact.validate(); err != nil {
		return nil, err
	}
	return &ValidatedOrderCustomerContact{
		OrderCustomerContact: *orderCustomerContact,
		isValidated:          false,
	}, nil
}
