package data

import (
	"github.com/materials-resources/s-prophet/infrastructure/data"
	"github.com/uptrace/bun"
)

type oeLine struct {
	data.OeLine `bun:",extend"`
	InvMast     *invMast `bun:"rel:has-one,join:inv_mast_uid=inv_mast_uid"`
}

type OeLineModel struct {
	bun bun.IDB
}
