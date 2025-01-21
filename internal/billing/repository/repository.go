package repository

import "github.com/uptrace/bun"

type Repository struct {
	Invoice *InvoiceRepository
}

func NewRepository(db bun.IDB) *Repository {
	return &Repository{
		Invoice: NewInvoiceRepository(db),
	}
}
