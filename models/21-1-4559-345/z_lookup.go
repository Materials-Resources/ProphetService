package gen

import "github.com/uptrace/bun"

type ZLookup struct {
	bun.BaseModel `bun:"table:z_lookup"`
	NValue        int32   `bun:"n_value,type:int,pk"`        // N Value
	ZValue        float64 `bun:"z_value,type:decimal(19,9)"` // Z Value
}
