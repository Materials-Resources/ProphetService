package prophet

import "github.com/uptrace/bun"

type TempSysobjects struct {
	bun.BaseModel `bun:"table:temp_sysobjects"`
	Name          string  `bun:"name,type:sysname"`
	Type          *string `bun:"type,type:char(2)"`
}
