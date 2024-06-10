package data

import "github.com/materials-resources/s-prophet/pkg/models"

type address struct {
	models.Address `bun:",extend"`
}
