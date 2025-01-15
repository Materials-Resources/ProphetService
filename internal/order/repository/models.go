package repository

import (
	"context"
	prophet "github.com/materials-resources/s-prophet/models/21-1-4559-345"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/schema"
	"time"
)

type address struct {
	prophet.Address `bun:",extend"`
}

type contacts struct {
	prophet.Contacts `bun:",extend"`
}

type customer struct {
	prophet.Customer `bun:",extend"`
}

type invMast struct {
	prophet.InvMast    `bun:",extend"`
	InvLocs            []*invLoc            `bun:"rel:has-many,join:inv_mast_uid=inv_mast_uid"`
	InventorySuppliers []*inventorySupplier `bun:"rel:has-many,join:inv_mast_uid=inv_mast_uid"`
}

type invLoc struct {
	prophet.InvLoc `bun:",extend"`

	InvMast *invMast `bun:"rel:belongs-to,join:inv_mast_uid=inv_mast_uid"`
}

type inventorySupplier struct {
	prophet.InventorySupplier `bun:",extend"`
}

type oeHdr struct {
	prophet.OeHdr `bun:",extend"`
	Customer      *customer `bun:"rel:has-one,join:customer_id=customer_id,join:company_id=company_id"`
	Contact       *contacts `bun:"rel:has-one,join:contact_id=id"`
	OeLines       []*oeLine `bun:"rel:has-many,join:oe_hdr_uid=oe_hdr_uid"`
}

var _ bun.BeforeAppendModelHook = (*oeHdr)(nil)

func (m *oeHdr) BeforeAppendModel(ctx context.Context, query bun.Query) error {
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

type oeHdrSalesrep struct {
	prophet.OeHdrSalesrep `bun:",extend"`
}

var _ bun.BeforeAppendModelHook = (*oeHdrSalesrep)(nil)

func (m *oeHdrSalesrep) BeforeAppendModel(ctx context.Context, query bun.Query) error {
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

type oeLine struct {
	prophet.OeLine `bun:",extend"`
	InvMast        *invMast `bun:"rel:has-one,join:inv_mast_uid=inv_mast_uid"`
}

var _ bun.BeforeAppendModelHook = (*oeLine)(nil)

func (m *oeLine) BeforeAppendModel(ctx context.Context, query schema.Query) error {
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

type oePickTicket struct {
	prophet.OePickTicket `bun:",extend"`
	OeHdr                *oeHdr   `bun:"rel:has-one,join:order_no=order_no"`
	CarrierAddress       *address `bun:"rel:has-one,join:carrier_id=id,join_on:carrier_flag='Y'"`
}

type quoteHdr struct {
	prophet.QuoteHdr `bun:",extend"`
}
