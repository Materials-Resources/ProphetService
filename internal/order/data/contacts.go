package data

import "github.com/materials-resources/s-prophet/infrastructure/data"

type contacts struct {
	data.Contacts `bun:",extend"`
}
