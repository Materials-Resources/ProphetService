package model

type TempSysobjects struct {
	bun.BaseModel `bun:"table:temp_sysobjects"`
	Name          `bun:"name,type:nvarchar(128)"`
	Type          string `bun:"type,type:char(2),nullzero"`
}
