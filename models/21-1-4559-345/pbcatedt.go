package gen

import "github.com/uptrace/bun"

type Pbcatedt struct {
	bun.BaseModel `bun:"table:pbcatedt"`
	PbeName       string `bun:"pbe_name,type:varchar(30),pk"`        // Powerbuilder catalog information
	PbeEdit       string `bun:"pbe_edit,type:varchar(254),nullzero"` // Powerbuilder catalog information
	PbeType       int16  `bun:"pbe_type,type:smallint"`              // Powerbuilder catalog information
	PbeCntr       int32  `bun:"pbe_cntr,type:int,nullzero"`          // Powerbuilder catalog information
	PbeSeqn       int16  `bun:"pbe_seqn,type:smallint,pk"`           // Powerbuilder catalog information
	PbeFlag       int32  `bun:"pbe_flag,type:int,nullzero"`          // Powerbuilder catalog information
	PbeWork       string `bun:"pbe_work,type:char(32),nullzero"`     // Powerbuilder catalog information
}
