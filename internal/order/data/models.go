package data

import "github.com/uptrace/bun"

type Models struct {
	OeHdr        OeHdrModel
	Customer     CustomerModel
	OePickTicket OePickTicketModel
}

func NewModels(bun bun.IDB) *Models {
	return &Models{
		OeHdr:        NewOeHdrModel(bun),
		Customer:     CustomerModel{bun: bun},
		OePickTicket: NewOePickTicketModel(bun),
	}

}
