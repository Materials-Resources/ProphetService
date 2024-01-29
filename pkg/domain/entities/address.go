package entities

import (
	"errors"
)

type Address struct {
	Name       string
	LineOne    string
	LineTwo    string
	City       string
	State      string
	PostalCode string
}

func (e *Address) validate() error {
	if e.Name == "" {
		return errors.New("invalid address details")
	}
	return nil
}

func NewAddress(name, lineOne, lineTwo, city, state, postalCode string) *Address {
	return &Address{
		Name:       name,
		LineOne:    lineOne,
		LineTwo:    lineTwo,
		City:       city,
		State:      state,
		PostalCode: postalCode,
	}
}
