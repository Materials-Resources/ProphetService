package model

import "github.com/uptrace/bun"

type Pbcatedt struct {
	bun.BaseModel `bun:"table:pbcatedt"`
	PbeName       string `bun:"pbe_name,type:varchar(30),pk"`
	PbeEdit       string `bun:"pbe_edit,type:varchar(254),nullzero"`
	PbeType       int16  `bun:"pbe_type,type:smallint"`
	PbeCntr       int32  `bun:"pbe_cntr,type:int,nullzero"`
	PbeSeqn       int16  `bun:"pbe_seqn,type:smallint,pk"`
	PbeFlag       int32  `bun:"pbe_flag,type:int,nullzero"`
	PbeWork       string `bun:"pbe_work,type:char(32),nullzero"`
}
