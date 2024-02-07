package domain

import (
	"errors"
)

type PickTicket struct {
	ID                 string
	ShippingAddress    ValidatedAddress
	OrderId            string
	OrderPurchaseOrder string
	OrderContactName   string
}

func (e *PickTicket) validate() error {
	if e.ID == "" {
		return errors.New("invalid pick ticket details")
	}
	return nil
}

func NewPickTicket() *PickTicket {
	return &PickTicket{}
}
