package repository

import (
	"errors"
	"github.com/uptrace/bun"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type Repository struct {
	Branch *BranchRepository
}

func NewRepository(db bun.IDB) *Repository {
	return &Repository{
		Branch: NewBranchRepository(db),
	}
}
