package data

import (
	prophet "github.com/materials-resources/s-prophet/models/21-1-4559-345"
	"github.com/uptrace/bun"
)

type InvoiceHdr struct {
	prophet.InvoiceHdr `bun:",extend"`
}

type Models struct {
	bun bun.IDB

	InvoiceHdr InvoiceHdrModel
}

func NewModels(bun bun.IDB) *Models {
	return &Models{
		bun: bun,

		InvoiceHdr: NewInvoiceHdrModel(bun),
	}
}
