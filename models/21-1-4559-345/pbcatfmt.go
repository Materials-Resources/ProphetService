package model

type Pbcatfmt struct {
	bun.BaseModel `bun:"table:pbcatfmt"`
	PbfName       string `bun:"pbf_name,type:varchar(30),pk"`
	PbfFrmt       string `bun:"pbf_frmt,type:varchar(254)"`
	PbfType       int16  `bun:"pbf_type,type:smallint"`
	PbfCntr       int32  `bun:"pbf_cntr,type:int,nullzero"`
}
