package data

import (
	"github.com/materials-resources/s-prophet/infrastructure/data"
	"github.com/uptrace/bun"
)

type customer struct {
	data.Customer `bun:",extend"`
}

type CustomerModel struct {
	bun bun.IDB
}
