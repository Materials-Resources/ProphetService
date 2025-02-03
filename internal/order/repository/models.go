package repository

import (
	"context"
	prophet "github.com/materials-resources/s-prophet/models/21-1-4559-345"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/schema"
	"time"
)

type Address struct {
	prophet.Address `bun:",extend"`
}

type Contacts struct {
	prophet.Contacts `bun:",extend"`
}

type Customer struct {
	prophet.Customer `bun:",extend"`
}

type InvMast struct {
	prophet.InvMast    `bun:",extend"`
	InvLocs            []*InvLoc            `bun:"rel:has-many,join:inv_mast_uid=inv_mast_uid"`
	InventorySuppliers []*InventorySupplier `bun:"rel:has-many,join:inv_mast_uid=inv_mast_uid"`
}

type InvLoc struct {
	prophet.InvLoc `bun:",extend"`

	InvMast *InvMast `bun:"rel:belongs-to,join:inv_mast_uid=inv_mast_uid"`
}

type InventorySupplier struct {
	prophet.InventorySupplier `bun:",extend"`
}

type NoteArea struct {
	prophet.NoteArea `bun:",extend"`
}

func (m NoteArea) BeforeAppendModel(ctx context.Context, query schema.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		m.DateCreated = time.Now()
		m.DateLastModified = time.Now()
		m.DeleteFlag = "N"
	case *bun.UpdateQuery:
		m.DateLastModified = time.Now()
	}
	return nil
}

var _ bun.BeforeAppendModelHook = (*NoteArea)(nil)

type OeHdr struct {
	prophet.OeHdr `bun:",extend"`
	Customer      *Customer       `bun:"rel:has-one,join:customer_id=customer_id,join:company_id=company_id"`
	Contact       *Contacts       `bun:"rel:has-one,join:contact_id=id"`
	OeLines       []*OeLine       `bun:"rel:has-many,join:oe_hdr_uid=oe_hdr_uid"`
	QuoteHdr      *QuoteHdr       `bun:"rel:has-one,join:oe_hdr_uid=oe_hdr_uid"`
	ShipTo        *ShipTo         `bun:"rel:has-one,join:address_id=ship_to_id,join:company_id=company_id,join:customer_id=customer_id"`
	OeHdrNotePads []*OeHdrNotepad `bun:"rel:has-many,join:order_no=order_no"`
}

var _ bun.BeforeAppendModelHook = (*OeHdr)(nil)

func (m *OeHdr) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		m.DateCreated = time.Now()
		m.DateLastModified = time.Now()
		m.DeleteFlag = "N"
	case *bun.UpdateQuery:
		m.DateLastModified = time.Now()
	}
	return nil
}

type OeHdrSalesrep struct {
	prophet.OeHdrSalesrep `bun:",extend"`
}

var _ bun.BeforeAppendModelHook = (*OeHdrSalesrep)(nil)

func (m *OeHdrSalesrep) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		m.DateCreated = time.Now()
		m.DateLastModified = time.Now()
		m.DeleteFlag = "N"
	case *bun.UpdateQuery:
		m.DateLastModified = time.Now()
	}
	return nil
}

type OeHdrNotepad struct {
	prophet.OeHdrNotepad `bun:",extend"`
	NoteArea             *NoteArea `bun:"rel:belongs-to,join:note_id=note_id"`
}

func (m OeHdrNotepad) BeforeAppendModel(ctx context.Context, query schema.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		m.DateCreated = time.Now()
		m.DateLastModified = time.Now()
		m.DeleteFlag = "N"
	case *bun.UpdateQuery:
		m.DateLastModified = time.Now()
	}
	return nil
}

var _ bun.BeforeAppendModelHook = (*OeHdrNotepad)(nil)

type OeLine struct {
	prophet.OeLine `bun:",extend"`
	InvMast        *InvMast `bun:"rel:belongs-to,join:inv_mast_uid=inv_mast_uid"`
	OeHdr          *OeHdr   `bun:"rel:belongs-to,join:oe_hdr_uid=oe_hdr_uid"`
}

var _ bun.BeforeAppendModelHook = (*OeLine)(nil)

func (m *OeLine) BeforeAppendModel(ctx context.Context, query schema.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		m.DateCreated = time.Now()
		m.DateLastModified = time.Now()
		m.DeleteFlag = "N"
	case *bun.UpdateQuery:
		m.DateLastModified = time.Now()
	}
	return nil
}

type OePickTicket struct {
	prophet.OePickTicket `bun:",extend"`
	OeHdr                *OeHdr   `bun:"rel:has-one,join:order_no=order_no"`
	CarrierAddress       *Address `bun:"rel:has-one,join:carrier_id=id,join_on:carrier_flag='Y'"`
}

type QuoteHdr struct {
	prophet.QuoteHdr `bun:",extend"`
}

func (m QuoteHdr) BeforeAppendModel(ctx context.Context, query schema.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		m.DateCreated = time.Now()
		m.DateLastModified = time.Now()
	case *bun.UpdateQuery:
		m.DateLastModified = time.Now()
	}
	return nil
}

var _ bun.BeforeAppendModelHook = (*QuoteHdr)(nil)

type ShipTo struct {
	prophet.ShipTo `bun:",extend"`
}
