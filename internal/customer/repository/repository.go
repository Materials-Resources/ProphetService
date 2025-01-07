package repository

import "github.com/uptrace/bun"

type Repository struct {
	Contacts *contactsRepository
}

func NewRepository(db *bun.DB) *Repository {
	return &Repository{
		Contacts: newContactsRepository(db),
	}
}
