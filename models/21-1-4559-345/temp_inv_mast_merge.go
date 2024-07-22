package prophet

import "github.com/uptrace/bun"

type TempInvMastMerge struct {
	bun.BaseModel        `bun:"table:temp_inv_mast_merge"`
	BadItemId            *string `bun:"bad_item_id,type:varchar(40)"`
	GoodItemId           *string `bun:"good_item_id,type:varchar(40)"`
	ProcessedFlag        *string `bun:"processed_flag,type:char(1)"`
	CreateAltCodeFlag    *string `bun:"create_alt_code_flag,type:char(1)"`
	KeepRestrictionsFlag *string `bun:"keep_restrictions_flag,type:char(1)"`
}
