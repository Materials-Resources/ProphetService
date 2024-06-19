package gen

import "github.com/uptrace/bun"

type P21Number struct {
	bun.BaseModel `bun:"table:p21_number"`
	Number        int16 `bun:"number,type:tinyint,pk"`
}
