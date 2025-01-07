package repository

import (
	"context"
	prophet "github.com/materials-resources/s-prophet/models/21-1-4559-345"
	"github.com/uptrace/bun"
)

type Contacts struct {
	prophet.Contacts `bun:",extend"`
	ContactsXShipTo  []*ContactsXShipTo `bun:"rel:has-many,join:id=contact_id"`
}

type ContactsXShipTo struct {
	prophet.ContactsXShipTo `bun:",extend"`
	Address                 Address `bun:"rel:has-one,join:ship_to_id=id"`
}

type contactsRepository struct {
	db *bun.DB
}

func newContactsRepository(db *bun.DB) *contactsRepository {
	return &contactsRepository{db: db}
}

func (r contactsRepository) Select(ctx context.Context, id string) (*Contacts, error) {
	var contact Contacts
	err := r.db.NewSelect().Model(&contact).Relation("ContactsXShipTo").Relation("ContactsXShipTo.Address").Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return &contact, nil

}
