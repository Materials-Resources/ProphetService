package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type TempCounterTable struct {
	bun.BaseModel  `bun:"table:temp_counter_table"`
	Uid            int32     `bun:"uid,type:int,autoincrement"`
	JobPriceHdrUid int32     `bun:"job_price_hdr_uid,type:int"`
	StartDate      time.Time `bun:"start_date,type:datetime"`
	EndDate        time.Time `bun:"end_date,type:datetime"`
}
