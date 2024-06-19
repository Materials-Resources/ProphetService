package gen

import "github.com/uptrace/bun"

type EccShortCodeDelete struct {
	bun.BaseModel         `bun:"table:ecc_short_code_delete"`
	EccShortCodeDeleteUid int32  `bun:"ecc_short_code_delete_uid,type:int,autoincrement,scanonly,pk"`
	InvMastUid            int32  `bun:"inv_mast_uid,type:int"`
	ShortCode             string `bun:"short_code,type:varchar(40)"`
}
