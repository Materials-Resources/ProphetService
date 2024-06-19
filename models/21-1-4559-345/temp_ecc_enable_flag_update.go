package gen

import "github.com/uptrace/bun"

type TempEccEnableFlagUpdate struct {
	bun.BaseModel              `bun:"table:temp_ecc_enable_flag_update"`
	TempEccEnableFlagUpdateUid int32 `bun:"temp_ecc_enable_flag_update_uid,type:int,autoincrement,identity"`
	InvMastUid                 int32 `bun:"inv_mast_uid,type:int"`
}
