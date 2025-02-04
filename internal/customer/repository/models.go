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

type InvMast struct {
	prophet.InvMast `bun:",extend"`
}

type InvoiceHdr struct {
	prophet.InvoiceHdr `bun:",extend"`

	InvoiceLines []*InvoiceLine `bun:"rel:has-many,join:invoice_no=invoice_no"`
}

type InvoiceLine struct {
	prophet.InvoiceLine `bun:",extend"`

	InvoiceHdr *InvoiceHdr `bun:"rel:belongs-to,join:invoice_no=invoice_no"`
}

type OeHdr struct {
	prophet.OeHdr `bun:",extend"`
	OeLines       []*OeLine `bun:"rel:has-many,join:oe_hdr_uid=oe_hdr_uid"`
}

type OeLine struct {
	prophet.OeLine `bun:",extend"`
	InvMast        *InvMast `bun:"rel:belongs-to,join:inv_mast_uid=inv_mast_uid"`
	OeHdr          *OeHdr   `bun:"rel:belongs-to,join:oe_hdr_uid=oe_hdr_uid"`
}

type ShipTo struct {
	prophet.ShipTo `bun:",extend"`
	Address        Address `bun:"rel:has-one,join:ship_to_id=id"`
}
