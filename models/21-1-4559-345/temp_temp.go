package model

type TempTemp struct {
	bun.BaseModel `bun:"table:temp_temp"`
	A             int32 `bun:"a,type:int"`
}
