package gen

import "github.com/uptrace/bun"

type Pbcatvld struct {
	bun.BaseModel `bun:"table:pbcatvld"`
	PbvName       string `bun:"pbv_name,type:varchar(30),pk"`       // Powerbuilder catalog information
	PbvVald       string `bun:"pbv_vald,type:varchar(254)"`         // Powerbuilder catalog information
	PbvType       int16  `bun:"pbv_type,type:smallint"`             // Powerbuilder catalog information
	PbvCntr       int32  `bun:"pbv_cntr,type:int,nullzero"`         // Powerbuilder catalog information
	PbvMsg        string `bun:"pbv_msg,type:varchar(254),nullzero"` // Powerbuilder catalog information
}
