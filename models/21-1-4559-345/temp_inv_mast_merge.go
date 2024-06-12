package model

import "github.com/uptrace/bun"

type TempInvMastMerge struct {
	bun.BaseModel        `bun:"table:temp_inv_mast_merge"`
	BadItemId            string `bun:"bad_item_id,type:varchar(40),nullzero"`
	GoodItemId           string `bun:"good_item_id,type:varchar(40),nullzero"`
	ProcessedFlag        string `bun:"processed_flag,type:char,nullzero"`
	CreateAltCodeFlag    string `bun:"create_alt_code_flag,type:char,nullzero"`
	KeepRestrictionsFlag string `bun:"keep_restrictions_flag,type:char,nullzero"`
}
