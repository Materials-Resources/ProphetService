package data

import (
	"errors"
	prophet "github.com/materials-resources/s-prophet/models/21-1-4559-345"
	"github.com/uptrace/bun"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type invoiceHdr struct {
	prophet.InvoiceHdr `bun:",extend"`
}

type Models struct {
	Invoices InvoiceModel
}

func NewModels(db bun.IDB) *Models {
	return &Models{
		Invoices: InvoiceModel{db: db},
	}
}
