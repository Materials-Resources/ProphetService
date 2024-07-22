package prophet

import "github.com/uptrace/bun"

type PcState struct {
	bun.BaseModel `bun:"table:pc_state"`
	StateSkey     int32   `bun:"state_skey,type:int,pk"`      // Padlock table
	CountrySkey   *int32  `bun:"country_skey,type:int"`       // Padlock table
	StateCode     string  `bun:"state_code,type:char(6)"`     // Padlock table
	StateName     *string `bun:"state_name,type:varchar(40)"` // Padlock table
}
