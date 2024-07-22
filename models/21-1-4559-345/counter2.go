package prophet

import "github.com/uptrace/bun"

type Counter2 struct {
	bun.BaseModel `bun:"table:counter2"`
	Id            string  `bun:"id,type:varchar(64)"`
	Description   *string `bun:"description,type:varchar(255)"`
	CounterNum    int32   `bun:"counter_num,type:int"`
}
