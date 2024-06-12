package model

import "github.com/uptrace/bun"

type TempTemp struct {
	bun.BaseModel `bun:"table:temp_temp"`
	A             int32 `bun:"a,type:int"`
}
