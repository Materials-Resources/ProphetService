package model

import "github.com/uptrace/bun"

type TempSysobjects struct {
	bun.BaseModel `bun:"table:temp_sysobjects"`
	Name          string `bun:"name,type:nvarchar(128)"`
	Type          string `bun:"type,type:char(2),nullzero"`
}
