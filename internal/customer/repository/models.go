package repository

import (
	prophet "github.com/materials-resources/s-prophet/models/21-1-4559-345"
)

type Address struct {
	prophet.Address `bun:",extend"`
}

type Contacts struct {
	prophet.Contacts `bun:",extend"`
	ContactsXShipTo  []*ContactsXShipTo `bun:"rel:has-many,join:id=contact_id"`
}

type ContactsXShipTo struct {
	prophet.ContactsXShipTo `bun:",extend"`
	ShipTo                  ShipTo `bun:"rel:belongs-to,join:ship_to_id=ship_to_id,join:company_id=company_id"`
}

type ShipTo struct {
	prophet.ShipTo `bun:",extend"`
	Address        Address `bun:"rel:has-one,join:ship_to_id=id"`
}
