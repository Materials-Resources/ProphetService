package data

import "github.com/uptrace/bun"

type Models struct {
	InvLoc       InvLocModel
	ProductGroup ProductGroupModel
}

func NewModels(bun bun.IDB) *Models {
	return &Models{
		InvLoc:       NewInvLocModel(bun),
		ProductGroup: NewProductGroupModel(bun),
	}
}
