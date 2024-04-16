package prophet_21_1_4559

import (
	"errors"
	"github.com/uptrace/bun"
)

var (
	ErrNotFound        = errors.New("could not find requested record")
	ErrDuplicateRecord = errors.New("a record with this key already exists")
)

type DeleteFlag int

const (
	Undefined DeleteFlag = iota
	DeleteFlagYes
	DeleteFlagNo
)

func (f DeleteFlag) String() string {
	return [...]string{"", "Y", "N"}[f]
}

type Models struct {
	InvMast      InvMastModel
	InvLoc       InvLocModel
	OeHdr        OeHdrModel
	ProductGroup ProductGroupModel
}

func NewModels(db *bun.DB) *Models {
	return &Models{
		InvMast: InvMastModel{bun: db},
		InvLoc:  InvLocModel{bun: db}, ProductGroup: ProductGroupModel{bun: db}, OeHdr: OeHdrModel{bun: db},
	}

}
