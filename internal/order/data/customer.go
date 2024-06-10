package data

import (
	"github.com/materials-resources/s-prophet/pkg/models"
	"github.com/uptrace/bun"
)

type customer struct {
	models.Customer `bun:",extend"`
}

type CustomerModel struct {
	bun bun.IDB
}
