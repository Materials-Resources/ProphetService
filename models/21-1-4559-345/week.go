package model

type Week struct {
	bun.BaseModel `bun:"table:week"`
	WeekUid       int32 `bun:"week_uid,type:int,pk,identity"`
	WeekNo        int32 `bun:"week_no,type:int"`
}
