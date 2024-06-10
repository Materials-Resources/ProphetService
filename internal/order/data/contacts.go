package data

import "github.com/materials-resources/s-prophet/pkg/models"

type contacts struct {
	models.Contacts `bun:",extend"`
}
