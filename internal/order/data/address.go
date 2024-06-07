package data

import "github.com/materials-resources/s-prophet/infrastructure/data"

type address struct {
	data.Address `bun:",extend"`
}
