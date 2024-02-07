package domain

import (
	"errors"
)

type Order struct {
	ID                string
	ShippingAddressId string
	ShippingAddress   ValidatedAddress
	CustomerContact   ValidatedOrderCustomerContact
	CustomerId        string
	CustomerName      string
	Items             []ValidatedOrderItem
	PurchaseOrder     string
}

func (o *Order) validate() error {
	if o.ID == "" {
		return errors.New("invalid order details")
	}
	return nil
}

func NewOrder(customerContactId string, purchaseOrder string) *Order {
	return &Order{}
}

type OrderItem struct {
	ID              string
	SN              string
	Name            string
	QuantityOrdered float64
}

func (p *OrderItem) validate() error {
	return nil
}

func NewOrderItem() *OrderItem {
	return &OrderItem{}
}

type OrderCustomerContact struct {
	Id          string
	Name        string
	PhoneNumber string
	Title       string
}

func (e *OrderCustomerContact) validate() error { return nil }
