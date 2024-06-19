package gen

import "github.com/uptrace/bun"

type Week struct {
	bun.BaseModel `bun:"table:week"`
	WeekUid       int32 `bun:"week_uid,type:int,autoincrement,identity,pk"`
	WeekNo        int32 `bun:"week_no,type:int"` // Week Number
}
