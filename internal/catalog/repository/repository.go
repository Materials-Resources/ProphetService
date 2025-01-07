package repository

import "github.com/uptrace/bun"

type Repository struct {
	ProductGroup *ProductGroupRepository
	Product      *ProductRepository
	InvMast      invMastRepository
}

func NewRepository(db bun.IDB) *Repository {
	return &Repository{ProductGroup: NewProductGroupRepository(db), InvMast: newInvMastRepository(db)}
}
