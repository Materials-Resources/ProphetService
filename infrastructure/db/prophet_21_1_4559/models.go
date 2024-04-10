package prophet_21_1_4559

import (
	"errors"
	"github.com/uptrace/bun"
)

var (
	ErrNotFound = errors.New("could not find requested record")
)

type Models struct {
	InvLoc InvLocModel
}

func NewModels(db *bun.DB) *Models {
	return &Models{
		InvLoc: InvLocModel{bun: db},
	}

}
