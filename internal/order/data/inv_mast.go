package data

import "github.com/materials-resources/s-prophet/infrastructure/data"

type invMast struct {
	data.InvMast `bun:",extend"`
}
