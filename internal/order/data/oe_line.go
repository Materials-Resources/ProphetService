package data

import (
	"github.com/materials-resources/s-prophet/pkg/models"
	"github.com/uptrace/bun"
)

type oeLine struct {
	models.OeLine `bun:",extend"`
	InvMast       *invMast `bun:"rel:has-one,join:inv_mast_uid=inv_mast_uid"`
}

type OeLineModel struct {
	bun bun.IDB
}
