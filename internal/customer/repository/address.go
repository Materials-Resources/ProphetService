package repository

import prophet "github.com/materials-resources/s-prophet/models/21-1-4559-345"

type Address struct {
	prophet.Address `bun:",extend"`
}
