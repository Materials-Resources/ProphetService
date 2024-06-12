package model

import "github.com/uptrace/bun"

type Pbcatvld struct {
	bun.BaseModel `bun:"table:pbcatvld"`
	PbvName       string `bun:"pbv_name,type:varchar(30),pk"`
	PbvVald       string `bun:"pbv_vald,type:varchar(254)"`
	PbvType       int16  `bun:"pbv_type,type:smallint"`
	PbvCntr       int32  `bun:"pbv_cntr,type:int,nullzero"`
	PbvMsg        string `bun:"pbv_msg,type:varchar(254),nullzero"`
}
