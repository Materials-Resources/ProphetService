package model

import "github.com/uptrace/bun"

type PcState struct {
	bun.BaseModel `bun:"table:pc_state"`
	StateSkey     int32  `bun:"state_skey,type:int,pk"`
	CountrySkey   int32  `bun:"country_skey,type:int,nullzero"`
	StateCode     string `bun:"state_code,type:char(6)"`
	StateName     string `bun:"state_name,type:varchar(40),nullzero"`
}
