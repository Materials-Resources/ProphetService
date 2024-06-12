package model

import "github.com/uptrace/bun"

type PcTimeZone struct {
	bun.BaseModel `bun:"table:pc_time_zone"`
	TimeZoneSkey  int32   `bun:"time_zone_skey,type:int,pk"`
	Description   string  `bun:"description,type:varchar(100)"`
	GmtOffsetHrs  float64 `bun:"gmt_offset_hrs,type:float"`
}
