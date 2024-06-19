package gen

import "github.com/uptrace/bun"

type Pbcatfmt struct {
	bun.BaseModel `bun:"table:pbcatfmt"`
	PbfName       string `bun:"pbf_name,type:varchar(30),pk"` // Powerbuilder catalog information
	PbfFrmt       string `bun:"pbf_frmt,type:varchar(254)"`   // Powerbuilder catalog information
	PbfType       int16  `bun:"pbf_type,type:smallint"`       // Powerbuilder catalog information
	PbfCntr       int32  `bun:"pbf_cntr,type:int,nullzero"`   // Powerbuilder catalog information
}
