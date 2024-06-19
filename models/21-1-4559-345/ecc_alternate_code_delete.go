package gen

import "github.com/uptrace/bun"

type EccAlternateCodeDelete struct {
	bun.BaseModel             `bun:"table:ecc_alternate_code_delete"`
	EccAlternateCodeDeleteUid int32  `bun:"ecc_alternate_code_delete_uid,type:int,autoincrement,pk"`
	AlternateCodeUid          int32  `bun:"alternate_code_uid,type:int"`
	InvMastUid                int32  `bun:"inv_mast_uid,type:int"`
	AlternateCode             string `bun:"alternate_code,type:varchar(40)"`
}
