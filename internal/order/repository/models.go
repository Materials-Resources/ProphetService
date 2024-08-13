package repository

import (
	"context"
	prophet "github.com/materials-resources/s-prophet/models/21-1-4559-345"
	"github.com/uptrace/bun"
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
	prophet.InvMast `bun:",extend"`
}

type oeHdr struct {
	prophet.OeHdr `bun:",extend"`
	Customer      *customer `bun:"rel:has-one,join:customer_id=customer_id,join:company_id=company_id"`
	Contact       *contacts `bun:"rel:has-one,join:contact_id=id"`
	OeLines       []*oeLine `bun:"rel:has-many,join:order_no=order_no"`
}

var _ bun.BeforeAppendModelHook = (*oeHdr)(nil)

func (m *oeHdr) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		m.DateCreated = time.Now()
	case *bun.UpdateQuery:
		m.DateLastModified = time.Now()
	}
	return nil
}

type oeLine struct {
	prophet.OeLine `bun:",extend"`
	InvMast        *invMast `bun:"rel:has-one,join:inv_mast_uid=inv_mast_uid"`
}

type oePickTicket struct {
	prophet.OePickTicket `bun:",extend"`
	OeHdr                *oeHdr   `bun:"rel:has-one,join:order_no=order_no"`
	CarrierAddress       *address `bun:"rel:has-one,join:carrier_id=id,join_on:carrier_flag='Y'"`
}
