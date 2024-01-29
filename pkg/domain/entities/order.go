package entities

import (
	"errors"
)

type Order struct {
	ID              string
	ShippingAddress ValidatedAddress
	CustomerContact struct {
		Id          string
		Name        string
		PhoneNumber string
		Email       string
	}
	Items         []ValidatedOrderItem
	PurchaseOrder string
}

func (o *Order) validate() error {
	if o.ID == "" {
		return errors.New("invalid order details")
	}
	return nil
}

func NewOrder(customerId string) *Order {
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
