package data

import "github.com/uptrace/bun"

type Models struct {
	InvLoc       InvLocModel
	InvMast      InvMastModel
	ProductGroup ProductGroupModel
}

func NewModels(bun bun.IDB) *Models {
	return &Models{
		InvLoc:       NewInvLocModel(bun),
		InvMast:      NewInvMastModel(bun),
		ProductGroup: NewProductGroupModel(bun),
	}
}
