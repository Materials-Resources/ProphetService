package model

import (
	"github.com/uptrace/bun"
)

type Contact struct {
	bun.BaseModel `bun:"table:contacts"`
	Id            string `bun:"id,type:varchar(16),pk"`
	FirstName     string `bun:"first_name,type:varchar(15)"`
	LastName      string `bun:"last_name,type:varchar(24)"`
}
